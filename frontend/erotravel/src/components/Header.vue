<template>
  <header class="flex justify-between border-b border-slate-100 px-10 py-8">
    <div class="flex items-center gap-4">
      <img src="/rasskazer.png" alt="Logo" class="w-20 rounded-3xl" />
      <div>
        <h2 class="text-xl font-bold uppercase">Рассказер.ру</h2>
        <p class="text-slate-400">Аудиокниги и фанфики</p>
      </div>
    </div>

    <ul class="flex items-center gap-10">
      <!--      <li class="flex items-center cursor-pointer gap-3 text-gray-500 hover:text-black">-->
      <!--        <img src="/cart.svg" alt="Cart" /><b>1205 руб.</b>-->
      <!--      </li>-->

      <li class="flex items-center cursor-pointer gap-2 text-gray-500 hover:text-black">
        <img src="/heart.svg" alt="Favorites" /><span>Закладки</span>
      </li>

      <li class="flex items-center cursor-pointer gap-2 text-gray-500 hover:text-black">
        <img src="/profile.svg" alt="Profile" /><span>Профиль</span>
      </li>
    </ul>
  </header>
</template>

<script setup></script>
