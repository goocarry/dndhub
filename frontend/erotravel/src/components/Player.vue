<template>
  <div>
    <div class="border border-slate-100 p-4 rounded-xl">
      <img src="/filo.webp" alt="Image" class="w-40" />

      <div class="flex justify-between">
        <p>Теоретический минимум по Computer Science</p>
      </div>
    </div>

    <div class="mt-8 border border-slate-100 p-8 pb-4 rounded-xl">
      <div class="border border-slate-500"></div>
      <div class="flex justify-between items-center px-4 pt-4">
        <button class="transition hover:-translate-y-1">
          <img src="/heart.svg" alt="Favorite" />
        </button>
        <button class="transition hover:-translate-y-1">
          <img src="/play.svg" alt="Play" class="w-10" />
        </button>
        <button class="transition hover:-translate-y-1">
          <img src="/search.svg" alt="Search" class="w-5" />
        </button>
      </div>
    </div>
  </div>
</template>
<script setup></script>
