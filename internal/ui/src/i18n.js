import { addMessages, init, getLocaleFromNavigator, locale as $locale, _ } from 'svelte-i18n'
import en from './locales/en.json'
import de from './locales/de.json'

addMessages('en', en)
addMessages('de', de)

init({
  fallbackLocale: 'en',
  initialLocale: getLocaleFromNavigator(),
})

export { $locale as locale, _ }

