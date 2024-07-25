import chunk from 'lodash/chunk'

import { withTelemetry } from '@shared-ui/common/services/opentelemetry'
import { fetchApi, security } from '@shared-ui/common/services'
import { FetchApiReturnType } from '@shared-ui/common/types/API.types'

import { OpenidConfigurationReturnType } from '@/containers/ApiTokens/ApiTokens.types'
import { ApiTokensApiEndpoints } from '@/containers/ApiTokens/constants'
import { SecurityConfig } from '@/containers/App/App.types'

export const getApiTokenUrlApi: () => Promise<string> = () => {
    const { m2mOauthClient } = security.getWellKnownConfig()

    return new Promise((resolve, reject) => {
        withTelemetry(
            () =>
                fetchApi(`${m2mOauthClient.authority}/.well-known/openid-configuration`, {
                    useToken: false,
                }),
            'get-m2mOauthClient-wellKnow-configuration'
        )
            .then((result: FetchApiReturnType<OpenidConfigurationReturnType>) => {
                if (result.data) {
                    resolve(result?.data?.token_endpoint.split('m2m-oauth-server')[0])
                } else {
                    reject(false)
                }
            })
            .catch((error: any) => {
                reject(error)
            })
    })
}

export const createApiTokenApi = async (body: any) => {
    const { cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig

    return getApiTokenUrlApi().then((result) => {
        const url = result.endsWith('/') ? result.slice(0, -1) : result
        return withTelemetry(
            () =>
                fetchApi(`${url}${ApiTokensApiEndpoints.API_TOKENS}`, {
                    method: 'POST',
                    body,
                    cancelRequestDeadlineTimeout,
                }),
            'create-api-token'
        )
    })
}

export const removeApiTokenApi = (ids: string[], chunkSize = 50) => {
    const { cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig

    const chunks = chunk(ids, chunkSize)
    return Promise.all(
        chunks.map((ids) => {
            return getApiTokenUrlApi().then((result) => {
                const url = result.endsWith('/') ? result.slice(0, -1) : result
                // const idsString = ids.map((id) => `${filterName}=${id}`).join('&')

                return withTelemetry(
                    () =>
                        fetchApi(`${url}${ApiTokensApiEndpoints.API_TOKENS_BLACKLIST}`, {
                            method: 'POST',
                            cancelRequestDeadlineTimeout,
                            body: {
                                idFilter: ids,
                            },
                        }),
                    `api-tokens-blacklist-${ids.join('-')}`
                )
            })
        })
    )
}
