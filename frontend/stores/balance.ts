import { defineStore } from 'pinia'

interface BalanceResponse {
  data: {
    balance: number
  }
}

interface BalanceState {
  amount: number | null
  loading: boolean
  error: string | null
}

export const useBalanceStore = defineStore('balance', {
  state: (): BalanceState => ({
    amount: null,
    loading: false,
    error: null
  }),

  actions: {
    async fetchCurrentBalance(): Promise<void> {
      this.loading = true
      this.error = null

      try {
        const response = await $fetch<BalanceResponse>(
          '/api/balances/current',
          {
            credentials: 'include'
          }
        )

        this.amount = response.data.balance

      } catch (error: unknown) {
        if (error instanceof Error) {
          this.error = error.message
        } else {
          this.error = 'Balance could not be fetched'
        }
      } finally {
        this.loading = false
      }
    }
  }
})
