import { FC, useEffect, useState } from 'react'
import { useIntl } from 'react-intl'
import { toast } from 'react-toastify'

import { showSuccessToast } from '@shared-ui/components/new/Toast'
import { getApiErrorMessage } from '@shared-ui/common/utils'
import { useIsMounted } from '@shared-ui/common/hooks'
import { Emitter } from '@shared-ui/common/services/emitter'
import { PendingCommandsExpandableList } from '@/containers/PendingCommands'
import { messages as menuT } from '@shared-ui/components/new/Menu/Menu.i18n'

import { DEVICES_REGISTERED_UNREGISTERED_COUNT_EVENT_KEY, RESET_COUNTER } from '../../constants'
import { useDevicesList } from '../../hooks'
import { DevicesList } from '../DevicesList/DevicesList'
import DevicesListHeader from '../DevicesListHeader/DevicesListHeader'
import { deleteDevicesApi } from '../../rest'
import { handleDeleteDevicesErrors, sleep } from '../../utils'
import { messages as t } from '../../Devices.i18n'
import PageLayout from '@shared-ui/components/new/PageLayout'
import { DeleteModal } from '@shared-ui/components/new/Modal'

const DevicesListPage: FC<any> = () => {
    const { formatMessage: _ } = useIntl()
    const {
        data,
        loading,
        error: deviceError,
        refresh,
    }: {
        data: any
        loading: boolean
        error: any
        refresh: () => void
    } = useDevicesList()
    const [deleteModalOpen, setDeleteModalOpen] = useState(false)
    const [isAllSelected, setIsAllSelected] = useState(false)
    const [selectedDevices, setSelectedDevices] = useState([])
    const [singleDevice, setSingleDevice] = useState<null | string>(null)
    const [deleting, setDeleting] = useState(false)
    const [unselectRowsToken, setUnselectRowsToken] = useState(1)
    const isMounted = useIsMounted()
    const combinedSelectedDevices = singleDevice ? [singleDevice] : selectedDevices

    useEffect(() => {
        if (deviceError) {
            toast.error(getApiErrorMessage(deviceError))
        }
    }, [deviceError])

    const handleOpenDeleteModal = (deviceId?: string) => {
        if (typeof deviceId === 'string') {
            setSingleDevice(deviceId)
        }

        setDeleteModalOpen(true)
    }

    const handleCloseDeleteModal = () => {
        setSingleDevice(null)
        setDeleteModalOpen(false)
    }

    const handleRefresh = () => {
        refresh()

        // Unselect all rows from the table
        setUnselectRowsToken((prevValue) => prevValue + 1)

        // Reset the counter on the Refresh button
        Emitter.emit(DEVICES_REGISTERED_UNREGISTERED_COUNT_EVENT_KEY, RESET_COUNTER)
    }

    const deleteDevices = async () => {
        try {
            setDeleting(true)
            await deleteDevicesApi(combinedSelectedDevices)
            await sleep(200)

            if (isMounted.current) {
                showSuccessToast({
                    title: _(t.devicesDeleted),
                    message: _(t.devicesDeletedMessage),
                })

                setDeleting(false)
                setDeleteModalOpen(false)
                setSingleDevice(null)
                setSelectedDevices([])
                handleCloseDeleteModal()
                handleRefresh()
            }
        } catch (error) {
            setDeleting(false)
            handleDeleteDevicesErrors(error, _)
        }
    }

    const selectedDevicesCount = combinedSelectedDevices.length
    const selectedDeviceName = selectedDevicesCount === 1 && data ? data.find?.((d: any) => d.id === combinedSelectedDevices[0])?.name : null
    const loadingOrDeleting = loading || deleting

    return (
        <PageLayout
            breadcrumbs={[
                {
                    label: _(menuT.devices),
                },
            ]}
            header={<DevicesListHeader loading={loading} refresh={handleRefresh} />}
            loading={loading}
            title={_(menuT.devices)}
        >
            <DevicesList
                data={data}
                isAllSelected={isAllSelected}
                loading={loadingOrDeleting}
                onDeleteClick={handleOpenDeleteModal}
                selectedDevices={selectedDevices}
                setIsAllSelected={setIsAllSelected}
                setSelectedDevices={setSelectedDevices}
                unselectRowsToken={unselectRowsToken}
            />

            <PendingCommandsExpandableList />

            <DeleteModal
                onClose={handleCloseDeleteModal}
                footerActions={[
                    {
                        label: _(t.cancel),
                        onClick: handleCloseDeleteModal,
                        variant: 'tertiary',
                    },
                    {
                        label: _(t.delete),
                        onClick: deleteDevices,
                        variant: 'primary',
                    },
                ]}
                show={deleteModalOpen}
                title={_(t.deleteDeviceMessage)}
                subTitle={_(t.deleteDeviceMessageSubTitle)}
                deleteInformation={
                    selectedDevicesCount === 1
                        ? [
                              { label: _(t.deviceName), value: selectedDeviceName },
                              { label: _(t.deviceId), value: combinedSelectedDevices[0] },
                          ]
                        : []
                }
            />
        </PageLayout>
    )
}

DevicesListPage.displayName = 'DevicesListPage'

export default DevicesListPage
