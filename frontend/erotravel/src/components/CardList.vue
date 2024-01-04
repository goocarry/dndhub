<template>
  <div class="grid grid-cols-3 gap-5">
    <Card
      v-for="item in items"
      :key="item.id"
      :title="item.title"
      :image-url="item.cover_url"
      :author="item.author"
      :is-favorite="true"
      :on-click-favorite="onClickFavorite"
      :on-click-play="onClickPlay"
    />
  </div>
</template>

<script setup>
import Card from '@/components/Card.vue'

defineProps({
  items: Array
})
const onClickFavorite = () => {
  alert('Added!')
}
const onClickPlay = () => {
  alert('Play!')
}
</script>
