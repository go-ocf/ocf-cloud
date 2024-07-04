import { fetchApi, security } from '@shared-ui/common/services'
import { deleteByChunks } from '@shared-ui/common/services/api-utils'
import { withTelemetry } from '@shared-ui/common/services/opentelemetry'
import { FetchApiReturnType } from '@shared-ui/common/types/API.types'

import { SecurityConfig } from '@/containers/App/App.types'
import { DELETE_CHUNK_SIZE, SnippetServiceApiEndpoints } from '@/containers/SnippetService/constants'

const getWellKnow = () => security.getWellKnownConfig()

export const createConfigurationApi = (body: any) => {
    const { httpGatewayAddress, cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const url = getWellKnow()?.ui?.snippetService || httpGatewayAddress

    return withTelemetry(
        () =>
            fetchApi(`${url}${SnippetServiceApiEndpoints.CONFIGURATIONS}`, {
                method: 'POST',
                cancelRequestDeadlineTimeout,
                body,
            }),
        'create-resource-config'
    )
}

export const updateResourceConfigApi = (id: string, body: any) => {
    const { httpGatewayAddress, cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const url = getWellKnow()?.ui?.snippetService || httpGatewayAddress

    return withTelemetry(
        () =>
            fetchApi(`${url}${SnippetServiceApiEndpoints.CONFIGURATIONS}/${id}`, {
                method: 'PUT',
                cancelRequestDeadlineTimeout,
                body,
            }),
        `update-resource-config-${id}`
    )
}

export const invokeConfigurationApi = (id: string, body: any) => {
    const { httpGatewayAddress, cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const url = getWellKnow()?.ui?.snippetService || httpGatewayAddress

    return withTelemetry(
        () =>
            fetchApi(`${url}${SnippetServiceApiEndpoints.CONFIGURATIONS}/${id}`, {
                method: 'POST',
                cancelRequestDeadlineTimeout,
                body,
            }),
        `invoke-configurations-${id}-${body.deviceId}`
    )
}

export const createConditionApi = (body: any) => {
    const { httpGatewayAddress, cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const url = getWellKnow()?.ui?.snippetService || httpGatewayAddress

    return withTelemetry(
        () =>
            fetchApi(`${url}${SnippetServiceApiEndpoints.CONDITIONS}`, {
                method: 'POST',
                cancelRequestDeadlineTimeout,
                body,
            }),
        'create-condition'
    )
}

export const updateConditionApi = (id: string, body: any) => {
    const { httpGatewayAddress, cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const url = getWellKnow()?.ui?.snippetService || httpGatewayAddress

    return withTelemetry(
        () =>
            fetchApi(`${url}${SnippetServiceApiEndpoints.CONDITIONS}/${id}`, {
                method: 'PUT',
                cancelRequestDeadlineTimeout,
                body,
            }),
        `update-condition-${id}`
    )
}

export const deleteConfigurationsApi = (ids: string[]) => {
    const { httpGatewayAddress, cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const url = getWellKnow()?.ui?.snippetService || httpGatewayAddress

    return deleteByChunks(
        `${url}${SnippetServiceApiEndpoints.CONFIGURATIONS}`,
        ids,
        cancelRequestDeadlineTimeout,
        'snippet-service-delete-configurations',
        'httpIdFilter',
        '/all',
        DELETE_CHUNK_SIZE
    )
}

export const deleteConditionsApi = (ids: string[]) => {
    const { httpGatewayAddress, cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const url = getWellKnow()?.ui?.snippetService || httpGatewayAddress

    return deleteByChunks(
        `${url}${SnippetServiceApiEndpoints.CONDITIONS}`,
        ids,
        cancelRequestDeadlineTimeout,
        'snippet-service-delete-conditions',
        'httpIdFilter',
        '/all',
        DELETE_CHUNK_SIZE
    )
}

export const deleteAppliedConfigurationApi = (ids: string[]) => {
    const { httpGatewayAddress, cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const url = getWellKnow()?.ui?.snippetService || httpGatewayAddress

    return deleteByChunks(
        `${url}${SnippetServiceApiEndpoints.CONFIGURATIONS_APPLIED}`,
        ids,
        cancelRequestDeadlineTimeout,
        'snippet-service-delete-applied-devices-config',
        'idFilter',
        '',
        DELETE_CHUNK_SIZE
    )
}

export const getOauthToken: () => Promise<string> = async () => {
    const { cancelRequestDeadlineTimeout } = security.getGeneralConfig() as SecurityConfig
    const { unauthorizedCallback, m2mOauthClient } = security.getWellKnownConfig()

    return new Promise((resolve, reject) => {
        withTelemetry(
            () =>
                fetchApi(`${m2mOauthClient.authority}/.well-known/openid-configuration`, {
                    useToken: false,
                }),
            'get-m2mOauthClient-wellKnow-configuration'
        )
            .then((result: FetchApiReturnType<{ issuer: string; jwks_uri: string; token_endpoint: string }>) => {
                if (result.data) {
                    withTelemetry(
                        () =>
                            fetchApi(result?.data?.token_endpoint, {
                                method: 'POST',
                                cancelRequestDeadlineTimeout,
                                unauthorizedCallback,
                                body: {
                                    client_assertion: security.getAccessToken(),
                                    client_assertion_type: 'urn:ietf:params:oauth:client-assertion-type:jwt-bearer', // hardcoded
                                    client_id: m2mOauthClient.clientId,
                                    grant_type: 'client_credentials', // hardcoded
                                },
                            }),
                        'get-m2mOauthClient-token'
                    ).then((result: FetchApiReturnType<{ access_token: string; scope: string; token_type: string }>) => {
                        resolve(result.data?.access_token || '')
                    })
                }
            })
            .catch((error: any) => {
                console.error(error)
            })
    })
}
