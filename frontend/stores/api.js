// stores/api.ts
import { defineStore } from 'pinia'

export const useApiStore = defineStore('api', {
  state: () => ({
    apiKey: process.env.NUXT_PUBLIC_API_KEY
  })
})
