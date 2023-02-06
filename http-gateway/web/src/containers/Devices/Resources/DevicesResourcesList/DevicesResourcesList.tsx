import { FC, useMemo } from 'react'
import { useIntl } from 'react-intl'
import classNames from 'classnames'

import Badge from '@shared-ui/components/new/Badge'
import Table from '@shared-ui/components/new/Table'
import DevicesResourcesActionButton from '../DevicesResourcesActionButton'
import { RESOURCES_DEFAULT_PAGE_SIZE, devicesStatuses } from '../../constants'
import { messages as t } from '../../Devices.i18n'
import { Props } from './DevicesResourcesList.types'

const DevicesResourcesList: FC<Props> = ({
  data,
  onUpdate,
  onCreate,
  onDelete,
  deviceStatus,
  loading,
}) => {
  const { formatMessage: _ } = useIntl()

  const isUnregistered = deviceStatus === devicesStatuses.UNREGISTERED
  const greyedOutClassName = classNames({ 'grayed-out': isUnregistered })

  const columns = useMemo(
    () => [
      {
        Header: _(t.href),
        accessor: 'href',
        Cell: ({ value, row }: { value: any; row: any }) => {
          const {
            original: { deviceId, href },
          } = row
          if (isUnregistered) {
            return <span>{value}</span>
          }
          return (
            <div className="tree-expander-container">
              <span
                className="link reveal-icon-on-hover"
                onClick={() => onUpdate({ deviceId, href })}
              >
                {value}
              </span>
              <i className="fas fa-pen" />
            </div>
          )
        },
        style: { width: '100%' },
      },
      {
        Header: _(t.types),
        accessor: 'resourceTypes',
        Cell: ({ value }: { value: any }) => {
          return (
            <div className="badges-box-horizontal">
              {value?.map?.((type: string) => (
                <Badge key={type}>{type}</Badge>
              ))}
            </div>
          )
        },
      },
      {
        Header: _(t.actions),
        accessor: 'actions',
        disableSortBy: true,
        Cell: ({ row }: { row: any }) => {
          const {
            original: { deviceId, href, interfaces },
          } = row
          return (
            <DevicesResourcesActionButton
              disabled={isUnregistered || loading}
              href={href}
              deviceId={deviceId}
              interfaces={interfaces}
              onCreate={onCreate}
              onUpdate={onUpdate}
              onDelete={onDelete}
            />
          )
        },
        className: 'actions',
      },
    ],
    [onUpdate, onCreate, onDelete, isUnregistered, loading] //eslint-disable-line
  )

  return (
    <Table
      columns={columns}
      data={data || []}
      defaultSortBy={[
        {
          id: 'href',
          desc: false,
        },
      ]}
      defaultPageSize={RESOURCES_DEFAULT_PAGE_SIZE}
      autoFillEmptyRows
      className={greyedOutClassName}
      paginationProps={{
        className: greyedOutClassName,
        disabled: isUnregistered,
      }}
    />
  )
}

DevicesResourcesList.displayName = 'DevicesResourcesList'

export default DevicesResourcesList