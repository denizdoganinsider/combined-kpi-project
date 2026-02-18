import { defineStore } from 'pinia'
import { useRuntimeConfig } from '#imports'

interface ApiState {
  apiKey: string
}

export const useApiStore = defineStore('api', {
  state: (): ApiState => ({
    apiKey: ''
  }),

  actions: {
    init(): void {
      const config = useRuntimeConfig()
      this.apiKey = config.public.apiKey as string
    }
  }
})
