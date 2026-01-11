import { defineStore } from 'pinia'

export interface Transaction {
  id: number
  fromUser: number
  toUser?: number
  amount: number
  type: 'credit' | 'debit' | 'transfer'
  status: 'pending' | 'completed' | 'failed'
  createdAt: string
}

export const useTransactionsStore = defineStore('transactions', {
  state: () => ({
    transactions: [] as Transaction[],
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchHistory(userId: number) {
      this.loading = true
      this.error = null

      try {
        const data = await $fetch<Transaction[]>(
          `/api/transactions/history/${userId}`
        )
        this.transactions = data
      } catch (err: any) {
        this.error = err?.data?.message || 'Failed to load transactions'
      } finally {
        this.loading = false
      }
    },

    async credit(userId: number, amount: number) {
      return await $fetch('/api/transactions/credit', {
        method: 'POST',
        body: { user_id: userId, amount }
      })
    },

    async debit(userId: number, amount: number) {
      return await $fetch('/api/transactions/debit', {
        method: 'POST',
        body: { user_id: userId, amount }
      })
    },

    async transfer(fromUserId: number, toUserId: number, amount: number) {
      return await $fetch('/api/transactions/transfer', {
        method: 'POST',
        body: {
          from_user_id: fromUserId,
          to_user_id: toUserId,
          amount
        }
      })
    }
  }
})
