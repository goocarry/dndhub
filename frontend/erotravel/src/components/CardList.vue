<template>
  <div class="grid grid-cols-3 gap-5">
    <Card
      image-url="/filo.webp"
      title="Теоретический минимум по Computer Science"
      author="Владстон Феррейра Фило"
      :is-favorite="true"
      :on-click-favorite="onClickFavorite"
      :on-click-play="onClickPlay"
    />
    <Card
      image-url="/ozon.png"
      title="Сириус 6Б"
      author="ozon671games"
      :is-favorite="false"
      :on-click-favorite="onClickFavorite"
      :on-click-play="onClickPlay"
    />
    <Card
      image-url="/afanasiev.jpg"
      title="Русские заветные сказки"
      author="Афанасьев А.Н."
      :is-favorite="true"
      :on-click-favorite="onClickFavorite"
      :on-click-play="onClickPlay"
    />
  </div>
</template>

<script setup>
import Card from '@/components/Card.vue'

const onClickFavorite = () => {
  alert('Added!')
}
const onClickPlay = () => {
  alert('Play!')
}
</script>
