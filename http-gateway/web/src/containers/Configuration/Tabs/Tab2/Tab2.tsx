import React, { FC, useCallback, useContext, useEffect, useState } from 'react'
import { Controller, SubmitHandler, useForm } from 'react-hook-form'
import { useIntl } from 'react-intl'
import ReactDOM from 'react-dom'

import { colors } from '@shared-ui/components/Atomic/_utils/colors'
import { getThemeTemplate } from '@shared-ui/components/Atomic/_theme/template'
import Editor from '@shared-ui/components/Atomic/Editor'
import SimpleStripTable from '@shared-ui/components/Atomic/SimpleStripTable'
import BottomPanel from '@shared-ui/components/Layout/BottomPanel/BottomPanel'
import Button from '@shared-ui/components/Atomic/Button'
import { Row } from '@shared-ui/components/Atomic/SimpleStripTable/SimpleStripTable.types'
import Spacer from '@shared-ui/components/Atomic/Spacer'
import { useIsMounted } from '@shared-ui/common/hooks'
import AppContext from '@shared-ui/app/share/AppContext'
import FormGroup from '@shared-ui/components/Atomic/FormGroup'
import FormInput, { inputAligns } from '@shared-ui/components/Atomic/FormInput'

import { Props, Inputs } from './Tab2.types'
import { messages as t } from '../../ConfigurationPage.i18n'
import { messages as g } from '@/containers/Global.i18n'

const Tab2: FC<Props> = (props) => {
    const { isTabActive, resetForm } = props
    const { formatMessage: _ } = useIntl()
    const { collapsed } = useContext(AppContext)
    const isMounted = useIsMounted()

    const [loading, setLoading] = useState(false)

    const {
        handleSubmit,
        formState: { errors },
        getValues,
        reset,
        control,
        register,
        setValue,
    } = useForm<Inputs>({
        mode: 'all',
        reValidateMode: 'onSubmit',
        values: {
            themeName: 'custom theme',
            colorPalette: colors ?? {},
            logoHeight: 0,
            logoWidth: 0,
            logoSource: '',
        },
    })

    useEffect(() => {
        if (resetForm) {
            reset()
        }
    }, [reset, resetForm])

    const getBase64 = useCallback(
        (file: any) =>
            new Promise((resolve) => {
                let baseURL: any = ''
                // Make new FileReader
                let reader = new FileReader()

                // Convert the file to base64 text
                reader.readAsDataURL(file)

                // on reader load something...
                reader.onload = () => {
                    // Make a fileInfo Object
                    baseURL = reader.result
                    resolve(baseURL)
                }
            }),
        []
    )

    const handleFileInput = (e: any) => {
        getBase64(e.target.files[0]).then((result) => {
            if (typeof result === 'string') {
                setValue('logoSource', result)
            }
        })
    }

    const rows: Row[] = [
        {
            attribute: _(t.themeName),
            value: (
                <FormGroup error={errors.themeName ? _(t.themeNameError) : undefined} errorTooltip={true} fullSize={true} id='theme-name' marginBottom={false}>
                    <FormInput
                        inlineStyle
                        align={inputAligns.RIGHT}
                        placeholder={_(t.themeName)}
                        {...register('themeName', { required: true, validate: (val) => val !== '' })}
                    />
                </FormGroup>
            ),
        },
        {
            attribute: _(t.colorPalette),
            autoHeight: true,
            value: (
                <Controller
                    control={control}
                    name='colorPalette'
                    render={({ field: { onChange, value } }) => (
                        <Spacer style={{ width: '100%' }} type='py-4'>
                            <Editor height='500px' json={value} onChange={(data) => onChange(JSON.parse(data))} />
                        </Spacer>
                    )}
                />
            ),
        },
        {
            attribute: _(t.logoSource),
            value: <input name='file' onChange={handleFileInput} type='file' />,
        },
        {
            attribute: _(t.logoHeight),
            value: (
                <FormGroup
                    error={errors.logoHeight ? _(t.logoHeightError) : undefined}
                    errorTooltip={true}
                    fullSize={true}
                    id='logo-height'
                    marginBottom={false}
                >
                    <FormInput
                        inlineStyle
                        align={inputAligns.RIGHT}
                        placeholder={_(t.logoHeight)}
                        type='number'
                        {...register('logoHeight', { required: true, valueAsNumber: true })}
                    />
                </FormGroup>
            ),
        },
        {
            attribute: _(t.logoWidth),
            value: (
                <FormGroup error={errors.logoWidth ? _(t.logoWidthError) : undefined} errorTooltip={true} fullSize={true} id='logo-width' marginBottom={false}>
                    <FormInput
                        inlineStyle
                        align={inputAligns.RIGHT}
                        placeholder={_(t.logoWidth)}
                        type='number'
                        {...register('logoWidth', { required: true, valueAsNumber: true })}
                    />
                </FormGroup>
            ),
        },
    ]

    const onSubmit: SubmitHandler<Inputs> = () => {
        const values = getValues()
        setLoading(true)
        const customThemeName = values.themeName.replace(/\s+/g, '_').toLowerCase()

        const fileName = `${customThemeName}.json`
        const data = new Blob(
            [
                JSON.stringify({
                    [customThemeName]: getThemeTemplate(values.colorPalette, {
                        height: `${values.logoHeight}px`,
                        width: `${values.logoWidth}px`,
                        source: values.logoSource,
                    }),
                }),
            ],
            { type: 'text/json' }
        )
        const jsonURL = window.URL.createObjectURL(data)
        const link = document.createElement('a')
        document.body.appendChild(link)
        link.href = jsonURL
        link.setAttribute('download', fileName)
        link.click()
        document.body.removeChild(link)
        setLoading(false)
    }

    return (
        <div>
            <form onSubmit={handleSubmit(onSubmit)}>
                <SimpleStripTable rows={rows} />
            </form>
            {isMounted &&
                document.querySelector('#modal-root') &&
                ReactDOM.createPortal(
                    <BottomPanel
                        actionPrimary={
                            <Button
                                disabled={Object.keys(errors).length > 0}
                                loading={loading}
                                loadingText={_(g.loading)}
                                onClick={() => onSubmit(getValues())}
                                variant='primary'
                            >
                                {_(t.generate)}
                            </Button>
                        }
                        actionSecondary={
                            <Button disabled={loading} onClick={() => reset()} variant='secondary'>
                                {_(t.reset)}
                            </Button>
                        }
                        leftPanelCollapsed={collapsed}
                        show={isTabActive}
                    />,
                    document.querySelector('#modal-root') as Element
                )}
        </div>
    )
}

Tab2.displayName = 'Tab2'

export default Tab2
