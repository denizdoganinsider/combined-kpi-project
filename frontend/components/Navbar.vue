<template>
  <nav class="bg-gradient-to-r from-[#032541] to-[#052d49] text-white h-14 flex items-center px-6 justify-between">
    <div class="flex items-center space-x-4">
      <div class="text-2xl font-bold text-green-400">TMDB</div>
      <ul class="hidden md:flex space-x-6 text-sm font-semibold">
        <li><a href="#" class="hover:text-cyan-400">Movies</a></li>
        <li><a href="#" class="hover:text-cyan-400">TV Shows</a></li>
        <li><a href="#" class="hover:text-cyan-400">People</a></li>
        <li><a href="#" class="hover:text-cyan-400">More</a></li>
      </ul>
    </div>

    <div class="flex items-center space-x-4 text-sm font-semibold">
      <button class="text-lg hover:text-cyan-400">+</button>
      <button class="border px-2 py-0.5 rounded hover:border-cyan-400">EN</button>
      <a href="#" class="hover:text-cyan-400">Login</a>
      <a href="#" class="hover:text-cyan-400">Join TMDB</a>
      <button class="hover:text-cyan-400">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none"
             viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round"
                d="M21 21l-4.35-4.35m0 0A7.5 7.5 0 1110.5 3a7.5 7.5 0 016.15 13.65z" />
        </svg>
      </button>
    </div>
  </nav>
</template>

<script setup>

</script>
