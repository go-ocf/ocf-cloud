import React, { FC, useContext } from 'react'
import { useIntl } from 'react-intl'

import Headline from '@shared-ui/components/Atomic/Headline'
import Loadable from '@shared-ui/components/Atomic/Loadable'
import Spacer from '@shared-ui/components/Atomic/Spacer'
import SimpleStripTable from '@shared-ui/components/Atomic/SimpleStripTable'
import FormInput from '@shared-ui/components/Atomic/FormInput'
import FormGroup from '@shared-ui/components/Atomic/FormGroup'
import { FormContext } from '@shared-ui/common/context/FormContext'
import { useForm } from '@shared-ui/common/hooks'

import { messages as t } from '@/containers/DeviceProvisioning/LinkedHubs/LinkedHubs.i18n'
import { messages as g } from '@/containers/Global.i18n'
import { Props, Inputs } from './TabContent1.types'

const TabContent1: FC<Props> = (props) => {
    const { defaultFormData, loading } = props

    const { formatMessage: _ } = useIntl()
    const { updateData, setFormError, commonFormGroupProps, commonInputProps } = useContext(FormContext)

    const {
        formState: { errors },
        register,
    } = useForm<Inputs>({ defaultFormData, updateData, setFormError, errorKey: 'tab3Content1' })

    return (
        <form>
            <Headline type='h5'>{_(t.general)}</Headline>
            <Spacer type='pt-4'>
                <Loadable condition={!loading}>
                    <SimpleStripTable
                        rows={[
                            {
                                attribute: _(t.ownerClaim),
                                value: (
                                    <FormGroup
                                        {...commonFormGroupProps}
                                        error={errors.authorization?.ownerClaim ? _(g.requiredField, { field: _(t.ownerClaim) }) : undefined}
                                        id='authorization?.ownerClaim'
                                    >
                                        <FormInput
                                            {...commonInputProps}
                                            {...register('authorization.ownerClaim', {
                                                required: true,
                                                validate: (val) => val !== '',
                                            })}
                                            placeholder={_(t.ownerClaim)}
                                        />
                                    </FormGroup>
                                ),
                            },
                        ]}
                    />
                </Loadable>
            </Spacer>
        </form>
    )
}

TabContent1.displayName = 'TabContent1'

export default TabContent1
