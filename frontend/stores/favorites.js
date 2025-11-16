import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useFavoritesStore = defineStore('favorites', () => {
  const favorites = ref([])

  const fetchFavorites = async () => {
    try {
      const res = await fetch('/api/favorites')
      const data = await res.json()
      favorites.value = data.map(f => f.movie_id)
    } catch (e) {
      console.error('Failed to fetch favorites', e)
    }
  }

  const addFavorite = async (movieId) => {
    try {
      await fetch('/api/favorites', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ movie_id: movieId }),
      })
      if (!favorites.value.includes(movieId)) favorites.value.push(movieId)
    } catch (e) {
      console.error('Failed to add favorite', e)
    }
  }

  const removeFavorite = async (movieId) => {
    try {
      await fetch(`/api/favorites/${movieId}`, {
        method: 'DELETE',
      })
      favorites.value = favorites.value.filter(id => id !== movieId)
    } catch (e) {
      console.error('Failed to remove favorite', e)
    }
  }

  const isFavorited = (movieId) => favorites.value.includes(movieId)

  return { favorites, fetchFavorites, addFavorite, removeFavorite, isFavorited }
})
