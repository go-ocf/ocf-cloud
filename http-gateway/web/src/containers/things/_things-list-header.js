import { useState } from 'react'
import { useIntl } from 'react-intl'
import { useSelector, useDispatch } from 'react-redux'
import PropTypes from 'prop-types'

import { useEmitter } from '@/common/hooks'
import { Button } from '@/components/button'
import { Switch } from '@/components/switch'
import { ProvisionNewDevice } from './provision-new-device'
import {
  THINGS_STATUS_WS_KEY,
  THINGS_REGISTERED_UNREGISTERED_COUNT_EVENT_KEY,
  RESET_COUNTER,
} from './constants'
import { isNotificationActive, toggleActiveNotification } from './slice'
import { messages as t } from './things-i18n'

export const ThingsListHeader = ({ loading, refresh }) => {
  const { formatMessage: _ } = useIntl()
  const dispatch = useDispatch()
  const enabled = useSelector(isNotificationActive(THINGS_STATUS_WS_KEY))
  const [numberOfChanges, setNumberOfChanges] = useState(0)

  useEmitter(
    THINGS_REGISTERED_UNREGISTERED_COUNT_EVENT_KEY,
    numberOfNewChanges => {
      setNumberOfChanges(
        numberOfNewChanges === RESET_COUNTER
          ? 0
          : numberOfChanges + numberOfNewChanges
      )
    }
  )

  const refreshThings = () => {
    // Re-fetch the devices list
    refresh()
  }

  return (
    <div className="d-flex align-items-center">
      <ProvisionNewDevice />
      <Button
        disabled={numberOfChanges <= 0 || loading}
        onClick={refreshThings}
        className="m-r-30"
        icon="fa-sync"
      >
        {`${_(t.refresh)}`}
        <span className="m-l-5 yellow-circle">{`${
          numberOfChanges > 9 ? '9+' : numberOfChanges
        }`}</span>
      </Button>
      <Switch
        id="status-notifications"
        label={_(t.notifications)}
        checked={enabled}
        onChange={e => {
          if (e.target.checked) {
            // Request browser notifications
            // (browsers will explicitly disallow notification permission requests not triggered in response to a user gesture,
            // so we must call it to make sure the user has received a notification request)
            Notification?.requestPermission?.()
          }

          dispatch(toggleActiveNotification(THINGS_STATUS_WS_KEY))
        }}
      />
    </div>
  )
}

ThingsListHeader.propTypes = {
  loading: PropTypes.bool.isRequired,
  refresh: PropTypes.func.isRequired,
}
