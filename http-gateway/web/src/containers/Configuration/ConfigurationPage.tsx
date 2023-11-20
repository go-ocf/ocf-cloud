import React, { ReactNode, useContext, useMemo } from 'react'
import { Controller, SubmitHandler, useForm } from 'react-hook-form'
import { useIntl } from 'react-intl'
import { useSelector, useDispatch } from 'react-redux'
import ReactDOM from 'react-dom'

import BottomPanel from '@shared-ui/components/Layout/BottomPanel/BottomPanel'
import Button from '@shared-ui/components/Atomic/Button'
import AppContext from '@shared-ui/app/share/AppContext'
import Footer from '@shared-ui/components/Layout/Footer'
import PageLayout from '@shared-ui/components/Atomic/PageLayout'
import SimpleStripTable from '@shared-ui/components/Atomic/SimpleStripTable'
import FormSelect, { selectAligns } from '@shared-ui/components/Atomic/FormSelect'
import Notification from '@shared-ui/components/Atomic/Notification/Toast'
import { useIsMounted } from '@shared-ui/common/hooks'

import { Inputs } from './ConfigurationPage.types'
import notificationId from '@/notificationId'
import { messages as t } from './Configuration.i18n'
import { CombinedStoreType } from '@/store/store'
import { setTheme } from '@/containers/App/slice'

interface RowsType {
    attribute: string
    value: ReactNode
}

const ConfigurationPage = () => {
    const { formatMessage: _ } = useIntl()
    const { collapsed } = useContext(AppContext)
    const isMounted = useIsMounted()
    const dispatch = useDispatch()

    const appStore = useSelector((state: CombinedStoreType) => state.app)

    const options = useMemo(() => appStore.configuration.themes.map((t) => ({ value: t, label: t })), [appStore.configuration.themes])

    const defTheme = useMemo(() => options.find((o) => o.value === appStore.configuration?.theme) || options[0], [appStore, options])

    const {
        handleSubmit,
        formState: { errors, isDirty, dirtyFields },
        getValues,
        reset,
        control,
    } = useForm<Inputs>({
        mode: 'all',
        reValidateMode: 'onSubmit',
        values: {
            theme: defTheme,
        },
    })

    const rows: RowsType[] = [
        {
            attribute: 'Theme',
            value: (
                <Controller
                    control={control}
                    name='theme'
                    render={({ field: { onChange, name, ref } }) => (
                        <FormSelect
                            inlineStyle
                            align={selectAligns.RIGHT}
                            defaultValue={defTheme}
                            name={name}
                            onChange={onChange}
                            options={options}
                            ref={ref}
                        />
                    )}
                />
            ),
        },
    ]

    const onSubmit: SubmitHandler<Inputs> = (data) => {
        dispatch(setTheme(data.theme.value))

        Notification.success(
            { title: _(t.configurationUpdated), message: _(t.configurationUpdatedMessage) },
            { notificationId: notificationId.HUB_CONFIGURATION_UPDATE }
        )
    }

    return (
        <PageLayout
            footer={<Footer paginationComponent={<div id='paginationPortalTarget'></div>} recentTasksPortal={<div id='recentTasksPortalTarget'></div>} />}
            title='App Configuration'
        >
            <form onSubmit={handleSubmit(onSubmit)}>
                <SimpleStripTable rows={rows} />
            </form>
            {isMounted &&
                document.querySelector('#modal-root') &&
                ReactDOM.createPortal(
                    <BottomPanel
                        actionPrimary={
                            <Button disabled={Object.keys(errors).length > 0} onClick={() => onSubmit(getValues())} variant='primary'>
                                {_(t.saveChanges)}
                            </Button>
                        }
                        actionSecondary={
                            <Button onClick={() => reset()} variant='secondary'>
                                {_(t.reset)}
                            </Button>
                        }
                        attribute={_(t.changesMade)}
                        leftPanelCollapsed={collapsed}
                        show={isDirty}
                        value={`${Object.keys(dirtyFields).length} ${Object.keys(dirtyFields).length > 1 ? _(t.settings) : _(t.setting)}`}
                    />,
                    document.querySelector('#modal-root') as Element
                )}
        </PageLayout>
    )
}

export default ConfigurationPage
