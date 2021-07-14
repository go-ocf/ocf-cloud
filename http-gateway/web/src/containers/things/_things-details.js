import { createElement, memo } from 'react'
import { useIntl } from 'react-intl'
import PropTypes from 'prop-types'
import classNames from 'classnames'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'

import { Badge } from '@/components/badge'
import { Label } from '@/components/label'
import { getValue } from '@/common/utils'

import { thingsStatuses } from './constants'
import { thingShape } from './shapes'
import { messages as t } from './things-i18n'

export const ThingsDetails = memo(({ data, loading }) => {
  const { formatMessage: _ } = useIntl()
  const deviceStatus = data?.metadata?.status?.value
  const isOnline = thingsStatuses.ONLINE === deviceStatus
  const isUnregistered = thingsStatuses.UNREGISTERED === deviceStatus
  const LabelWithLoading = p =>
    createElement(Label, {
      ...p,
      inline: true,
      className: classNames({
        shimmering: loading,
        'grayed-out': isUnregistered,
      }),
    })

  return (
    <Row>
      <Col>
        <LabelWithLoading title="ID">
          {getValue(data?.id)}
        </LabelWithLoading>
        <LabelWithLoading title={_(t.types)}>
          <div className="align-items-end badges-box-vertical">
            {data?.types?.map?.(type => (
              <Badge key={type}>{type}</Badge>
            )) || '-'}
          </div>
        </LabelWithLoading>
      </Col>
      <Col>
        <LabelWithLoading title={_(t.status)}>
          <Badge className={isOnline ? 'green' : 'red'}>
            {isOnline ? _(t.online) : _(t.offline)}
          </Badge>
        </LabelWithLoading>
      </Col>
    </Row>
  )
})

ThingsDetails.propTypes = {
  data: thingShape,
  loading: PropTypes.bool,
}

ThingsDetails.defaultProps = {
  data: null,
  loading: false,
}
