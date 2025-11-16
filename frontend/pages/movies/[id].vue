<template>
    <div class="movie-details p-4 max-w-4xl mx-auto">
      <div class="flex flex-col md:flex-row gap-6">
        <img
          :src="movie.poster_url"
          :alt="movie.title"
          class="w-full md:w-1/3 rounded-lg shadow-lg"
        />
  
        <div class="flex-1">
          <h1 class="text-3xl font-bold mb-2">{{ movie.title }}</h1>
          <p class="text-gray-600 mb-4">{{ movie.overview }}</p>
  
          <h2 class="text-xl font-semibold mb-2">Cast</h2>
          <ul class="list-disc list-inside">
            <li v-for="actor in movie.cast" :key="actor.id">{{ actor.name }} as {{ actor.character }}</li>
          </ul>
  
          <button
            @click="toggleFavorite"
            class="mt-4 px-4 py-2 rounded bg-yellow-400 hover:bg-yellow-500 text-white font-semibold"
          >
            {{ isFavorited ? "Favoriden Kaldır" : "Favoriye Ekle" }}
          </button>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from "vue";
  import { useRoute } from "vue-router";
  import { useFavoritesStore } from "@/stores/favorites";
  
  const route = useRoute();
  const favoritesStore = useFavoritesStore();
  
  const movie = ref({
    title: "",
    poster_url: "",
    overview: "",
    cast: [],
  });
  
  const isFavorited = ref(false);
  
  const fetchMovie = async () => {
    const res = await fetch(`/api/movies/${route.params.id}`);
    const data = await res.json();
    movie.value = data;
  
    isFavorited.value = favoritesStore.isFavorited(movie.value.id);
  };
  
  const toggleFavorite = () => {
    if (isFavorited.value) {
      favoritesStore.removeFavorite(movie.value.id);
    } else {
      favoritesStore.addFavorite(movie.value.id);
    }
    isFavorited.value = !isFavorited.value;
  };
  
  onMounted(() => {
    fetchMovie();
  });
  </script>
  
  <style scoped>

  </style>
  