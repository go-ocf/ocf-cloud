import { test, expect, Page } from '@playwright/test'
import testId from '../../../../src/testId'

test('snippet-service-configurations-list-open', async ({ page }) => {
    await page.goto('')
    await page.getByTestId(testId.menu.snippetService.link).click()
    await page.getByTestId(testId.menu.snippetService.configurations).click()

    await expect(page).toHaveTitle(/Configuraions | plgd Dashboard/)
    await expect(page).toHaveScreenshot({ fullPage: true, omitBackground: true })
})

const fillAddForm = async (page: Page) => {
    await page.getByTestId(testId.snippetService.configurations.addPage.form.name).fill('my-cfg-2')

    await page.getByTestId(testId.snippetService.configurations.addPage.form.addResourceButton).click()
    await expect(page.getByTestId(`${testId.snippetService.configurations.addPage.form.createResourceModal}-modal`)).toBeVisible()
    await page.getByTestId(`${testId.snippetService.configurations.addPage.form.createResourceModal}-input-href`).fill('/oc/con')
    await page.getByTestId(`${testId.snippetService.configurations.addPage.form.createResourceModal}-editor-input`).fill('123')

    await expect(page.getByTestId(`${testId.snippetService.configurations.addPage.form.createResourceModal}-confirm-button`)).toBeVisible()
    await page.getByTestId(`${testId.snippetService.configurations.addPage.form.createResourceModal}-confirm-button`).click()
    await expect(page.getByTestId(`${testId.snippetService.configurations.addPage.form.createResourceModal}-modal`)).not.toBeVisible()
}

test('add-configuration-reset', async ({ page }) => {
    await page.goto('')
    await page.getByTestId(testId.menu.snippetService.link).click()
    await page.getByTestId(testId.menu.snippetService.configurations).click()

    await page.getByTestId(testId.snippetService.configurations.list.addConfigurationButton).click()
    await expect(page).toHaveTitle(/Create new Configuration | plgd Dashboard/)
    await expect(page).toHaveScreenshot({ fullPage: true, omitBackground: true })

    fillAddForm(page)

    await expect(page.getByTestId(testId.snippetService.configurations.addPage.form.resourceTable)).toBeVisible()
    await expect(page.getByTestId(`${testId.snippetService.configurations.addPage.form.resourceTable}-row-0`)).toBeVisible()

    await page.getByTestId(testId.snippetService.configurations.addPage.form.resetButton).click()

    await expect(page.getByTestId(`${testId.snippetService.configurations.addPage.form.resourceTable}-row-0`)).not.toBeVisible()
    await expect(page.getByTestId(testId.snippetService.configurations.addPage.form.name)).toHaveValue('')

    await expect(page.getByTestId(testId.snippetService.configurations.addPage.form.resetButton)).toBeDisabled()
    await expect(page.getByTestId(testId.snippetService.configurations.addPage.form.addButton)).toBeDisabled()
})

test('add-configuration-save', async ({ page }) => {
    await page.goto('')
    await page.getByTestId(testId.menu.snippetService.link).click()
    await page.getByTestId(testId.menu.snippetService.configurations).click()
    await page.getByTestId(testId.snippetService.configurations.list.addConfigurationButton).click()

    fillAddForm(page)

    await expect(page.getByTestId(testId.snippetService.configurations.addPage.form.addButton)).not.toBeDisabled()
    await page.getByTestId(testId.snippetService.configurations.addPage.form.addButton).click()

    await expect(page).toHaveTitle(/my-cfg-2 | plgd Dashboard/)
})

test('list-invoke-modal', async ({ page }) => {
    await page.goto('')
    await page.getByTestId(testId.menu.snippetService.link).click()
    await page.getByTestId(testId.menu.snippetService.configurations).click()

    page.setViewportSize({ width: 1600, height: 720 })

    await expect(page.getByTestId(testId.snippetService.configurations.list.table)).toBeVisible()
    await expect(page.getByTestId(`${testId.snippetService.configurations.list.table}-row-0`)).toBeVisible()
    await expect(page.getByTestId(`${testId.snippetService.configurations.list.table}-row-0-invoke`)).toBeVisible()

    await page.getByTestId(`${testId.snippetService.configurations.list.table}-row-0-invoke`).click()
    await expect(page.getByTestId(testId.snippetService.configurations.list.invokeModal)).toBeVisible()

    // close and open
    await expect(page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-close`)).toBeVisible()
    await page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-close`).click()
    await expect(page.getByTestId(testId.snippetService.configurations.list.invokeModal)).not.toBeVisible()

    await page.getByTestId(`${testId.snippetService.configurations.list.table}-row-0-invoke`).click()
    await expect(page.getByTestId(testId.snippetService.configurations.list.invokeModal)).toBeVisible()

    await page.locator('#deviceId').focus()
    await expect(page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-select-input`)).toBeVisible()

    // select
    await page.locator('#deviceId').click()
    await page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-select-input`).fill('3aae0672-47f3-4498-78d4-b061e6105ccd')
    await page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-select-3aae0672-47f3-4498-78d4-b061e6105ccd`).click()

    await expect(page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-footer-reset`)).toBeVisible()
    await expect(page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-footer-done`)).toBeVisible()

    await page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-footer-done`).click()
    await page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-force-label`).click()

    await expect(page).toHaveScreenshot({ fullPage: true, omitBackground: true })

    await expect(page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-reset`)).toBeVisible()
    await expect(page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-invoke`)).toBeVisible()

    await expect(page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-reset`)).not.toBeDisabled()
    await expect(page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-invoke`)).not.toBeDisabled()

    await page.getByTestId(`${testId.snippetService.configurations.list.invokeModal}-invoke`).click()
})
