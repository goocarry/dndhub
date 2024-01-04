<template>
  <div
    class="relative bg-white border border-slate-100 rounded-3xl p-8 cursor-pointer hover:-translate-y-2 hover:shadow-xl transition"
  >
    <img
      :src="isFavorite ? '/like-2.svg' : '/like-1.svg'"
      alt="Like 1"
      class="absolute top-8 left-8"
      @click="onClickFavorite"
    />
    <img :src="imageUrl" alt="Sneaker" class="h-40" />

    <p class="mt-2 mb-2">{{ title }}</p>

    <div class="flex justify-between relative">
      <div class="flex flex-col">
        <span class="text-slate-500">Автор: </span>
        <span>{{ author }}</span>
      </div>
    </div>

    <div class="absolute bottom-4 right-4 w-8">
      <img @click="onClickPlay" src="/play.svg" alt="Play" />
    </div>
  </div>
</template>

<script setup>
defineProps({
  imageUrl: String,
  title: String,
  author: String,
  isFavorite: Boolean,
  onClickFavorite: Function,
  onClickPlay: Function
})

// const onClickPlay = () => {
//   alert('play function')
// }
</script>
