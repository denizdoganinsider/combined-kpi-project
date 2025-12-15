import { defineStore } from 'pinia'

export const useBalanceStore = defineStore('balance', {
  state: () => ({
    amount: null,
    loading: false,
    error: null
  }),

  actions: {
    async fetchCurrentBalance() {
      this.loading = true
      this.error = null

      try {
        const { data } = await $fetch('/api/v1/balances/current', {
          credentials: 'include'
        })

        this.amount = data.balance
      } catch (err) {
        this.error = 'Balance could not be fetched'
      } finally {
        this.loading = false
      }
    }
  }
})
