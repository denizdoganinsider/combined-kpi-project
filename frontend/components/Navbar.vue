<template>
  <nav class="bg-gradient-to-r from-[#032541] to-[#052d49] text-white h-14 flex items-center px-6 justify-between relative">

    <div class="flex items-center space-x-6">
      <div class="flex items-center text-2xl font-bold">
        <span class="text-green-400">TMDB</span>
        <div class="w-6 h-4 ml-1 bg-cyan-400 rounded-full"></div>
      </div>

      <ul class="hidden md:flex space-x-6 text-sm font-semibold">
        <li class="relative group">
          <a href="#" class="hover:text-cyan-400">Movies</a>
          <ul class="absolute top-full left-0 mt-2 w-40 bg-white text-black rounded shadow-md z-50
                     opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-opacity">
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Popular</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Now Playing</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Upcoming</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Top Rated</a></li>
          </ul>
        </li>

        <li class="relative group">
          <a href="#" class="hover:text-cyan-400">TV Shows</a>
          <ul class="absolute top-full left-0 mt-2 w-40 bg-white text-black rounded shadow-md z-50
                     opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-opacity">
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Popular</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Airing Today</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">On TV</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Top Rated</a></li>
          </ul>
        </li>

        <li class="relative group">
          <a href="#" class="hover:text-cyan-400">People</a>
          <ul class="absolute top-full left-0 mt-2 w-40 bg-white text-black rounded shadow-md z-50
                     opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-opacity">
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Popular People</a></li>
          </ul>
        </li>

        <li class="relative group">
          <a href="#" class="hover:text-cyan-400">More</a>
          <ul class="absolute top-full left-0 mt-2 w-48 bg-white text-black rounded shadow-md z-50
                     opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-opacity">
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Discussions</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Leaderboard</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">Support</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">API Documentation</a></li>
            <li><a href="#" class="block px-4 py-2 hover:bg-gray-100">API for Business</a></li>
          </ul>
        </li>
      </ul>
    </div>

    <div class="flex items-center space-x-4 text-sm font-semibold">
      <button class="text-white text-xl">＋</button>
      <button class="border border-white px-2 py-[2px] rounded text-xs">EN</button>
      <a href="#" class="hover:text-cyan-400">Login</a>
      <a href="#" class="hover:text-cyan-400">Join TMDB</a>
      <button class="text-cyan-400 text-xl">🔍</button>
    </div>
  </nav>
</template>

<script setup>
</script>