<template>
  <div class="fixed top-0 left-0 h-full w-full bg-black z-10 opacity-70"></div>
  <div class="bg-white w-96 h-full fixed right-0 top-0 z-20 p-8">
    <DrawerHead />
    <Player />
  </div>
</template>

<script setup>
import DrawerHead from '@/components/DrawerHead.vue'
import Player from '@/components/Player.vue'
</script>
