export default defineNuxtConfig({
  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
  ],

  ssr: true,

  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080',
      apiKey: process.env.NUXT_PUPLIC_TMDB_API_KEY || ''
    }
  },

  app: {
    head: {
      title: 'Movie Tracker',
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Movie and TV tracking application' }
      ]
    }
  }
})
