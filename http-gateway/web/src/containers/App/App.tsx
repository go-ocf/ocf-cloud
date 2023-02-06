import { hot } from 'react-hot-loader/root'
import { useContext, useState, useEffect } from 'react'
import { useIntl } from 'react-intl'
import PageLoader from '@shared-ui/components/new/PageLoader'
import { security } from '@/common/services/security' // @PM: shared
import { openTelemetry } from '@shared-ui/common/services/opentelemetry'
import { messages as t } from './App.i18n'
import { AppContext } from './AppContext'
import './app.scss'
import { getAppWellKnownConfiguration } from '@/containers/App/AppRest'
import AppInner from '@/containers/App/AppInner/AppInner'
import { AuthProvider, UserManager } from 'oidc-react'

const App = () => {
  const { formatMessage: _ } = useIntl()
  const [wellKnownConfig, setWellKnownConfig] = useState<any>(null)
  const [wellKnownConfigFetched, setWellKnownConfigFetched] = useState(false)
  const [configError, setConfigError] = useState<any>(null)

  openTelemetry.init('hub')

  useEffect(() => {
    if (!wellKnownConfig && !wellKnownConfigFetched) {
      const fetchWellKnownConfig = async () => {
        try {
          const { data: wellKnown } = await openTelemetry.withTelemetry(
            () =>
              getAppWellKnownConfiguration(
                process.env.REACT_APP_HTTP_WELL_NOW_CONFIGURATION_ADDRESS ||
                  window.location.origin
              ),
            'get-hub-configuration'
          )

          const { webOauthClient, deviceOauthClient, ...generalConfig } =
            wellKnown

          const clientId = webOauthClient?.clientId
          const httpGatewayAddress = wellKnown.httpGatewayAddress
          const authority = wellKnown.authority

          if (!clientId || !authority || !httpGatewayAddress) {
            throw new Error(
              'clientId, authority, audience and httpGatewayAddress must be set in webOauthClient of web_configuration.json'
            )
          } else {
            // Set the auth configurations
            security.setGeneralConfig(generalConfig)
            security.setWebOAuthConfig(webOauthClient)
            security.setDeviceOAuthConfig(deviceOauthClient)

            setWellKnownConfigFetched(true)
            setWellKnownConfig(wellKnown)
          }
        } catch (e) {
          setConfigError(
            new Error('Could not retrieve the well-known configuration.')
          )
        }
      }

      fetchWellKnownConfig()
    }
  }, [wellKnownConfig, wellKnownConfigFetched])

  // Render an error box with an auth error
  if (configError) {
    return (
      <div className="client-error-message">
        {`${_(t.authError)}: ${configError?.message}`}
      </div>
    )
  }

  // Placeholder loader while waiting for the auth status
  if (!wellKnownConfig) {
    return (
      <>
        <PageLoader className="auth-loader" loading />
        <div className="page-loading-text">{`${_(t.loading)}...`}</div>
      </>
    )
  }

  const oidcCommonSettings = {
    authority: wellKnownConfig.authority,
    scope: wellKnownConfig.webOauthClient.scopes.join?.(' '),
  }

  const onSignIn = async () => {
    window.location.hash = ''
    window.location.href = window.location.origin
  }

  return (
    <AuthProvider
      {...oidcCommonSettings}
      clientId={wellKnownConfig.webOauthClient.clientId}
      redirectUri={window.location.origin}
      onSignIn={onSignIn}
      automaticSilentRenew={true}
      userManager={
        new UserManager({
          ...oidcCommonSettings,
          client_id: wellKnownConfig.webOauthClient.clientId,
          redirect_uri: window.location.origin,
          extraQueryParams: {
            audience: wellKnownConfig.webOauthClient.audience || undefined,
          },
        })
      }
    >
      <AppInner
        wellKnownConfig={wellKnownConfig}
        openTelemetry={openTelemetry}
      />
    </AuthProvider>
  )
}

export const useAppConfig = () => useContext(AppContext)

export default hot(App)