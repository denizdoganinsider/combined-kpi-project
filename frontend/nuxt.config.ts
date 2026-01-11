export default defineNuxtConfig({
  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
  ],
  compatibilityDate: '2025-04-18',

  runtimeConfig: {
    public: {
      apiBase: 'http://localhost:8080'
    }
  }
})
