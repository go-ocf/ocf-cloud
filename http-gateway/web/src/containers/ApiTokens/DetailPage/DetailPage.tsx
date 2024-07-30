import React, { FC, useMemo } from 'react'
import { generatePath, useParams } from 'react-router-dom'
import { useIntl } from 'react-intl'
import chunk from 'lodash/chunk'

import SimpleStripTable from '@shared-ui/components/Atomic/SimpleStripTable'
import Loadable from '@shared-ui/components/Atomic/Loadable'
import Headline from '@shared-ui/components/Atomic/Headline'
import Spacer from '@shared-ui/components/Atomic/Spacer'
import { Column } from '@shared-ui/components/Atomic/Grid'
import Row from '@shared-ui/components/Atomic/Grid/Row'

import { useApiTokenDetail } from '@/containers/ApiTokens/hooks'
import DetailHeader from './DetailHeader'
import PageLayout from '@/containers/Common/PageLayout'
import { messages as t } from '../ApiTokens.i18n'
import { pages } from '@/routes'
import { messages as g } from '@/containers/Global.i18n'
import { formatDateVal } from '@/containers/PendingCommands/DateFormat'
import { getExpiration, parseClaimData } from '@/containers/ApiTokens/utils'

const DetailPage: FC<any> = () => {
    const { apiTokenId } = useParams()
    const { formatMessage: _, formatDate, formatTime } = useIntl()
    const { data, loading } = useApiTokenDetail(apiTokenId || '', !!apiTokenId)

    const breadcrumbs = useMemo(
        () => [{ label: _(t.apiTokens), link: generatePath(pages.API_TOKENS.LINK) }, { label: data?.name || '' }],
        // eslint-disable-next-line react-hooks/exhaustive-deps
        [data]
    )

    const claimsData = useMemo(
        () =>
            parseClaimData({
                data,
                hidden: ['issuedAt', 'version', 'expiration', 'name'],
                dateFormat: ['auth_time'],
                formatTime,
                formatDate,
            }),
        [data, formatDate, formatTime]
    )

    const chunks = useMemo(() => chunk(claimsData, 2), [claimsData])

    return (
        <PageLayout
            breadcrumbs={breadcrumbs}
            header={<DetailHeader id={apiTokenId!} loading={loading} name={data?.name} />}
            loading={loading}
            title={data?.name}
        >
            <Loadable condition={!!data && !loading}>
                <SimpleStripTable
                    leftColSize={6}
                    rightColSize={6}
                    rows={[
                        {
                            attribute: _(g.name),
                            value: data.name,
                        },
                        {
                            attribute: _(g.id),
                            value: data.id,
                        },
                        {
                            attribute: _(t.expiration),
                            value: data.expiration
                                ? getExpiration(data.expiration, formatDate, formatTime, {
                                      expiredText: (formatedDate) => _(t.expiredDate, { date: formatedDate }),
                                      expiresOn: (formatedDate) => _(t.expiresOn, { date: formatedDate }),
                                      noExpirationDate: _(t.noExpirationDate),
                                  })
                                : '-',
                        },
                        {
                            attribute: _(t.issuedAt),
                            value: data.issuedAt ? formatDateVal(new Date(data.issuedAt * 1000), formatDate, formatTime) : '',
                        },
                    ]}
                />
            </Loadable>
            <Loadable condition={!!data && !loading}>
                <>
                    <Spacer type='mt-8 mb-4'>
                        <Headline type='h5'>{_(t.tokenClaims)}</Headline>
                    </Spacer>
                    <Row>
                        {chunks.map((chunk, key) => (
                            <Column key={`chunk-col-${key === 0 ? 'left' : 'right'}`} xxl={6}>
                                <SimpleStripTable leftColSize={6} rightColSize={6} rows={chunk} />
                            </Column>
                        ))}
                    </Row>
                </>
            </Loadable>
        </PageLayout>
    )
}

DetailPage.displayName = 'DetailPage'

export default DetailPage
