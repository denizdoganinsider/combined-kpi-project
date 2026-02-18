import { defineStore } from 'pinia'
import { ref } from 'vue'

interface FavoriteItem {
  movie_id: number
}

export const useFavoritesStore = defineStore('favorites', () => {
  const favorites = ref<number[]>([])

  const fetchFavorites = async (): Promise<void> => {
    try {
      const data = await $fetch<FavoriteItem[]>('/api/favorites')

      favorites.value = data.map((f) => f.movie_id)
    } catch (error: unknown) {
      console.error('Failed to fetch favorites', error)
    }
  }

  const addFavorite = async (movieId: number): Promise<void> => {
    try {
      await $fetch('/api/favorites', {
        method: 'POST',
        body: { movie_id: movieId }
      })

      if (!favorites.value.includes(movieId)) {
        favorites.value.push(movieId)
      }

    } catch (error: unknown) {
      console.error('Failed to add favorite', error)
    }
  }

  const removeFavorite = async (movieId: number): Promise<void> => {
    try {
      await $fetch(`/api/favorites/${movieId}`, {
        method: 'DELETE'
      })

      favorites.value = favorites.value.filter(
        (id) => id !== movieId
      )

    } catch (error: unknown) {
      console.error('Failed to remove favorite', error)
    }
  }

  const isFavorited = (movieId: number): boolean => {
    return favorites.value.includes(movieId)
  }

  return {
    favorites,
    fetchFavorites,
    addFavorite,
    removeFavorite,
    isFavorited
  }
})