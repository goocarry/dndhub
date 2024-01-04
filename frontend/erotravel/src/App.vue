<script setup>
import Header from '@/components/Header.vue'
import CardList from '@/components/CardList.vue'
import Drawer from '@/components/Drawer.vue'

const items = [
  {
    id: '1',
    title: '1984',
    author_id: '1',
    author: 'Оруэлл Дж.',
    contents_id: '1',
    narrator: 'Narrator X',
    genre: 'Антиутопия',
    release_year: 1949,
    rating: 4.8,
    cover_url: 'https://www.bic-media.com/dmrs/cover.do?isbn=9783730609767&width=966'
  },
  {
    id: '2',
    title: 'Преступление и наказание',
    author_id: '2',
    author: 'Достоевский Ф.М.',
    contents_id: '2',
    narrator: 'Narrator Y',
    genre: 'Классика',
    release_year: 1813,
    rating: 4.8,
    cover_url: 'https://kniga.lv/wp-content/uploads/2022/02/cover1__w600-315-510x797.jpg'
  },
  {
    id: '3',
    title: 'Убить пересмешника',
    author_id: '3',
    author: 'Харпер Ли',
    contents_id: '3',
    narrator: 'Narrator Z',
    genre: 'Роман',
    release_year: 1960,
    rating: 4.7,
    cover_url:
      'https://img4.labirint.ru/rc/3cbf09cf313c7e910541c6c55f934579/363x561q80/books63/621296/cover.jpg?1686677406'
  },
  {
    id: '4',
    title: 'Теоретический минимум по Computer Science',
    author_id: '4',
    author: 'Владстон Феррейра Фило',
    contents_id: '4',
    narrator: 'Narrator Z',
    genre: 'Учебная литература',
    release_year: 2022,
    rating: 4.8,
    cover_url: 'https://kniga.lv/wp-content/uploads/2020/05/itd000000000924641_cover1.jpg'
  },
  {
    id: '5',
    title: 'Русские заветные сказки',
    author_id: '5',
    author: 'Афанасьев А.Н.',
    contents_id: '5',
    narrator: 'Narrator Z',
    genre: 'Сказка',
    release_year: 1960,
    rating: 4.7,
    cover_url:
      'https://s1.livelib.ru/boocover/1001391752/o/d3f3/Aleksandr_Afanasev__Russkie_zavetnye_skazki._Tolko_dlya_vzroslyh._Tolko_dlya_muz.jpeg'
  },
  {
    id: '6',
    title: 'Сириус 6Б',
    author_id: '6',
    author: 'ozon671games',
    contents_id: '6',
    narrator: 'ozon671games',
    genre: 'Фантастика',
    release_year: 2023,
    rating: 4.9,
    cover_url: '/ozon.png'
  }
]
</script>

<template>
  <!--  <Drawer />-->
  <div class="bg-white w-4/5 m-auto rounded-xl shadow-xl mt-14">
    <Header />

    <div class="p-10">
      <div class="flex justify-between">
        <h2 class="text-3xl font-bold mb-8">Все аудиокниги</h2>

        <div class="relative">
          <img class="absolute top-3 left-4" src="/search.svg" alt="Search" />
          <input
            class="border rounded-md py-2 pl-11 pr-4 outline-none focus:border-gray-400"
            placeholder="Поиск..."
          />
        </div>
      </div>

      <div class="mt-10">
        <CardList :items="items" />
      </div>
    </div>
  </div>
</template>
