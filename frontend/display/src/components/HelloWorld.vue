<template>
  <!-- <v-container class="fill-height">
    <v-responsive class="d-flex align-center text-center fill-height">
      <v-img height="300" src="@/assets/logo.svg" />

      <div class="text-body-2 font-weight-light mb-n1">Welcome to</div>

      <h1 class="text-h2 font-weight-bold">Vuetify</h1>

      <div class="py-14" />

      <v-row class="d-flex align-center justify-center">
        <v-col cols="auto">
          <v-btn
            href="https://vuetifyjs.com/components/all/"
            min-width="164"
            rel="noopener noreferrer"
            target="_blank"
            variant="text"
          >
            <v-icon icon="mdi-view-dashboard" size="large" start />

            Components
          </v-btn>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="primary"
            href="https://vuetifyjs.com/introduction/why-vuetify/#feature-guides"
            min-width="228"
            rel="noopener noreferrer"
            size="x-large"
            target="_blank"
            variant="flat"
          >
            <v-icon icon="mdi-speedometer" size="large" start />

            Get Started
          </v-btn>
        </v-col>

        <v-col cols="auto">
          <v-btn
            href="https://community.vuetifyjs.com/"
            min-width="164"
            rel="noopener noreferrer"
            target="_blank"
            variant="text"
          >
            <v-icon icon="mdi-account-group" size="large" start />

            Community
          </v-btn>
        </v-col>
      </v-row>
    </v-responsive>
  </v-container> -->
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
          <v-btn class="ma-4 bg-white" prepend-icon="mdi-dumbbell" variant="tonal"> Начать! </v-btn>
        </div>
      </v-col>
    </v-img>
  </div>
</template>

<script setup>
//
</script>
