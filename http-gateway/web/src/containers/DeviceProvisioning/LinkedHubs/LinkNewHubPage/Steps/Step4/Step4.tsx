import React, { FC, useContext } from 'react'
import { useIntl } from 'react-intl'
import { Controller } from 'react-hook-form'

import FormLabel from '@shared-ui/components/Atomic/FormLabel'
import FormInput from '@shared-ui/components/Atomic/FormInput'
import { useForm } from '@shared-ui/common/hooks'
import FormGroup from '@shared-ui/components/Atomic/FormGroup'
import { FormContext } from '@shared-ui/common/context/FormContext'
import TimeoutControl from '@shared-ui/components/Atomic/TimeoutControl'
import Spacer from '@shared-ui/components/Atomic/Spacer'
import * as commonStyles from '@shared-ui/components/Templates/FullPageWizard/FullPageWizardCommon.styles'
import StepButtons from '@shared-ui/components/Templates/FullPageWizard/StepButtons'

import { messages as t } from '@/containers/DeviceProvisioning/LinkedHubs/LinkedHubs.i18n'
import { messages as g } from '@/containers/Global.i18n'
import { Inputs, Props } from './Step4.types'
import SubStepTls from '../SubStepTls'
import FullPageWizard from '@shared-ui/components/Templates/FullPageWizard'

const Step4: FC<Props> = (props) => {
    const { defaultFormData, onSubmit } = props
    const { formatMessage: _ } = useIntl()
    const { updateData, setFormError, setStep } = useContext(FormContext)

    const {
        formState: { errors },
        register,
        control,
        updateField,
        watch,
        setValue,
    } = useForm<Inputs>({ defaultFormData, updateData, setFormError, errorKey: 'step4' })

    return (
        <form>
            <h1 css={commonStyles.headline}>{_(t.authorization)}</h1>

            <FullPageWizard.SubHeadline noBorder>{_(t.general)}</FullPageWizard.SubHeadline>
            <FullPageWizard.Description>{_(t.addLinkedHubAuthorizationGeneralDescription)}</FullPageWizard.Description>

            <FormGroup error={errors.authorization?.ownerClaim ? _(g.requiredField, { field: _(t.ownerClaim) }) : undefined} id='authorization.ownerClaim'>
                <FormLabel text={_(t.ownerClaim)} />
                <FormInput
                    {...register('authorization.ownerClaim', {
                        required: true,
                        validate: (val) => val !== '',
                    })}
                    onBlur={(e) => updateField('authorization.ownerClaim', e.target.value)}
                />
            </FormGroup>

            <h2 css={commonStyles.subHeadline}>{_(t.oAuthClient)}</h2>
            <FullPageWizard.Description>{_(t.addLinkedHubAuthorizationoAuthClientDescription)}</FullPageWizard.Description>

            <FormGroup error={errors.authorization?.provider?.name ? _(g.requiredField, { field: _(t.name) }) : undefined} id='authorization.provider.name'>
                <FormLabel text={_(t.name)} />
                <FormInput
                    {...register('authorization.provider.name', {
                        required: true,
                        validate: (val) => val !== '',
                    })}
                    onBlur={(e) => updateField('authorization.provider.name', e.target.value)}
                />
            </FormGroup>
            <FormGroup
                error={errors.authorization?.provider?.clientId ? _(g.requiredField, { field: _(t.clientId) }) : undefined}
                id='authorization.provider.clientId'
            >
                <FormLabel text={_(t.clientId)} />
                <FormInput
                    {...register('authorization.provider.clientId', {
                        required: true,
                        validate: (val) => val !== '',
                    })}
                    onBlur={(e) => updateField('authorization.provider.clientId', e.target.value)}
                />
            </FormGroup>
            <FormGroup
                error={errors.authorization?.provider?.clientSecret ? _(g.requiredField, { field: _(t.clientSecret) }) : undefined}
                id='authorization.provider.clientSecret'
            >
                <FormLabel text={_(t.clientSecret)} />
                <FormInput
                    {...register('authorization.provider.clientSecret', {
                        required: true,
                        validate: (val) => val !== '',
                    })}
                    onBlur={(e) => updateField('authorization.provider.clientSecret', e.target.value)}
                />
            </FormGroup>
            <FormGroup
                error={errors.authorization?.provider?.authority ? _(g.requiredField, { field: _(t.authority) }) : undefined}
                id='authorization.provider.authority'
            >
                <FormLabel text={_(t.authority)} />
                <FormInput
                    {...register('authorization.provider.authority', {
                        required: true,
                        validate: (val) => val !== '',
                    })}
                    onBlur={(e) => updateField('authorization.provider.authority', e.target.value)}
                />
            </FormGroup>
            <Controller
                control={control}
                name='authorization.provider.scopes'
                render={({ field: { onChange, value } }) => (
                    <FormGroup
                        error={errors.authorization?.provider?.authority ? _(g.requiredField, { field: _(t.scopes) }) : undefined}
                        id='authorization.provider.scopes'
                    >
                        <FormLabel text={_(t.scopes)} />
                        <FormInput
                            onBlur={(e) => updateField('authorization.provider.scopes', e.target.value.split(' '))}
                            onChange={(e) => onChange(e.target.value.split(' '))}
                            value={Array.isArray(value) ? value.join(' ') : value}
                        />
                    </FormGroup>
                )}
            />

            <SubStepTls
                control={control}
                prefix='authorization.provider.http.'
                setValue={(field: string, value: any) => {
                    // @ts-ignore
                    setValue(field, value)
                    updateField(field, value)
                }}
                updateField={updateField}
                watch={watch}
            />

            <Spacer type='pt-12'>
                <FullPageWizard.SubHeadline>{_(t.hTTP)}</FullPageWizard.SubHeadline>
                <FullPageWizard.Description>{_(t.addLinkedHubAuthorizationHttpDescription)}</FullPageWizard.Description>
            </Spacer>

            <FormGroup
                error={errors?.authorization?.provider?.http?.maxIdleConns ? _(g.requiredField, { field: _(t.maxIdleConnections) }) : undefined}
                id='authorization.provider.http.maxIdleConns'
            >
                <FormLabel text={_(t.maxIdleConnections)} />
                <FormInput
                    {...register('authorization.provider.http.maxIdleConns', {
                        required: true,
                        valueAsNumber: true,
                    })}
                    type='number'
                />
            </FormGroup>

            <FormGroup
                error={errors?.authorization?.provider?.http?.maxConnsPerHost ? _(g.requiredField, { field: _(t.maxConnectionsPerHost) }) : undefined}
                id='authorization.provider.http.maxConnsPerHost'
            >
                <FormLabel text={_(t.maxConnectionsPerHost)} />
                <FormInput
                    {...register('authorization.provider.http.maxConnsPerHost', {
                        required: true,
                        valueAsNumber: true,
                    })}
                    type='number'
                />
            </FormGroup>

            <FormGroup
                error={errors?.authorization?.provider?.http?.maxIdleConnsPerHost ? _(g.requiredField, { field: _(t.maxIdleConnectionsPerHost) }) : undefined}
                id='authorization.provider.http.maxIdleConnsPerHost'
            >
                <FormLabel text={_(t.maxIdleConnectionsPerHost)} />
                <FormInput
                    {...register('authorization.provider.http.maxIdleConnsPerHost', {
                        required: true,
                        valueAsNumber: true,
                    })}
                    type='number'
                />
            </FormGroup>

            <Controller
                control={control}
                name='authorization.provider.http.idleConnTimeout'
                render={({ field: { onChange, value } }) => (
                    <TimeoutControl
                        watchUnitChange
                        align='left'
                        defaultTtlValue={parseInt(value, 10)}
                        defaultValue={parseInt(value, 10)}
                        i18n={{
                            default: _(g.default),
                            duration: _(t.idleConnectionTimeout),
                            unit: _(g.metric),
                            placeholder: _(g.placeholder),
                        }}
                        onChange={(v) => onChange(v.toString())}
                        rightStyle={{
                            width: 150,
                        }}
                    />
                )}
            />

            <Spacer type='mt-5'>
                <Controller
                    control={control}
                    name='authorization.provider.http.timeout'
                    render={({ field: { onChange, value } }) => (
                        <TimeoutControl
                            watchUnitChange
                            align='left'
                            defaultTtlValue={parseInt(value, 10)}
                            defaultValue={parseInt(value, 10)}
                            i18n={{
                                default: _(g.default),
                                duration: _(g.timeout),
                                unit: _(g.metric),
                                placeholder: _(g.placeholder),
                            }}
                            onChange={(v) => onChange(v.toString())}
                            rightStyle={{
                                width: 150,
                            }}
                        />
                    )}
                />
            </Spacer>

            <StepButtons
                i18n={{
                    back: _(g.back),
                    continue: _(g.create),
                }}
                onClickBack={() => setStep?.(2)}
                onClickNext={() => onSubmit?.()}
            />
        </form>
    )
}

Step4.displayName = 'Step4'

export default Step4
