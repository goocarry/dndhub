<template>
  <div class="fill-height" style="max-height: 100vh">
    <v-img class="bg-white fill-height" cover src="@/assets/gym.jpg">
      <v-col class="flex-column justify-space-between" style="height: 20vh">
        <div>
          <p class="ma-4 text-h5 font-weight-black text-white">
            <v-icon class="mx-2" icon="mdi-dumbbell" />
            MarathonApp
          </p>
        </div>
        <div class="fill-height"></div>
        <div class="ma-6">
          <p class="ma-4 text-h4 font-weight-black text-white">Персональные тренировки</p>
          <p class="ma-4 text-h4 font-weight-black text-white">от Анатолия Бицепса</p>
        </div>
        <div class="ma-6 fill-width">
          <Timer />
        </div>
        <div class="ma-6 fill-width">
          <v-btn class="ma-4 bg-white" prepend-icon="mdi-dumbbell" variant="tonal"> Начать! </v-btn>
        </div>
      </v-col>
    </v-img>
  </div>
</template>

<script>
import Timer from './Timer.vue'
export default {
  components:{ Timer },
}
</script>
