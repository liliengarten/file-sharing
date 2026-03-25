<template>
  <div class="home-page">
    <!-- ── Navigation ── -->

    <!-- ── Filter chips ── -->
    <div class="filters-bar">
      <button
          v-for="filter in filters"
          :key="filter"
          class="filter-chip"
          :class="{ active: activeFilter === filter }"
          @click="activeFilter = filter"
      >
        {{ filter }}
      </button>
    </div>

    <!-- ── Masonry feed ── -->
    <div class="feed">
      <div class="masonry">
        <pin-card v-for="pin in allPins" :key="id" :pin="pin"></pin-card>
      </div>

      <div class="load-wrap">
        <button class="btn-load" :disabled="loading" @click="loadMore">
          {{ loading ? 'Загрузка...' : 'Загрузить ещё' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import {ref, computed} from 'vue'
import {RouterLink, useRoute} from 'vue-router'
import PinCard from "@/components/PinCard.vue";

const $route = useRoute()

// ── State ──
const searchQuery = ref('')
const activeFilter = ref('Всё')
const loading = ref(false)
const page = ref(1)

// ── Data ──
const navLinks = [
  {to: '/', label: 'Лента'},
  {to: '/explore', label: 'Обзор'},
  {to: '/following', label: 'Подписки'},
]

const filters = [
  'Всё', 'Интерьер', 'Архитектура', 'Мода',
  'Еда', 'Путешествия', 'Природа', 'Искусство',
  'Фотография', 'Минимализм', 'Типографика',
]

const allPins = ref([
  {
    id: 1,
    title: 'Японский минимализм в жилом интерьере',
    author: '@marinadiz',
    authorInitials: 'МД',
    likes: 234,
    height: 'h-lg',
    palette: 'pal-1'
  },
  {
    id: 2,
    title: 'Скандинавская палитра',
    author: '@annaN',
    authorInitials: 'АН',
    likes: 89,
    height: 'h-sm',
    palette: 'pal-5'
  },
  {
    id: 3,
    title: 'Студия с открытой планировкой — Осло',
    author: '@ivanarch',
    authorInitials: 'ИВ',
    likes: 412,
    height: 'h-xl',
    palette: 'pal-3'
  },
  {
    id: 4,
    title: 'Бетон и дерево',
    author: '@klara',
    authorInitials: 'КЛ',
    likes: 178,
    height: 'h-md',
    palette: 'pal-7'
  },
  {id: 5, title: 'Тихий угол', author: '@olgas', authorInitials: 'ОС', likes: 56, height: 'h-xs', palette: 'pal-2'},
  {
    id: 6,
    title: 'Вечер в розовых тонах — Токио',
    author: '@dm_photo',
    authorInitials: 'ДМ',
    likes: 321,
    height: 'h-lg',
    palette: 'pal-6'
  },
  {
    id: 7,
    title: 'Монохромная кухня',
    author: '@tkitchen',
    authorInitials: 'ТК',
    likes: 145,
    height: 'h-md',
    palette: 'pal-4'
  },
  {
    id: 8,
    title: 'Прогулка по Хоккайдо — зимний лес',
    author: '@natalytravel',
    authorInitials: 'НЕ',
    likes: 677,
    height: 'h-xl',
    palette: 'pal-8'
  },
  {
    id: 9,
    title: 'Шрифтовой плакат',
    author: '@roman.p',
    authorInitials: 'РП',
    likes: 203,
    height: 'h-sm',
    palette: 'pal-9'
  },
  {
    id: 10,
    title: 'Каменная терраса с видом на горы',
    author: '@evgarch',
    authorInitials: 'ЕА',
    likes: 388,
    height: 'h-md',
    palette: 'pal-10'
  },
  {
    id: 11,
    title: 'Минималистичный стол',
    author: '@vlad',
    authorInitials: 'ВЛ',
    likes: 44,
    height: 'h-xs',
    palette: 'pal-1'
  },
  {
    id: 12,
    title: 'Лофт с кирпичными стенами — Берлин',
    author: '@margo',
    authorInitials: 'МГ',
    likes: 512,
    height: 'h-lg',
    palette: 'pal-3'
  },
])

// ── Computed ──
const visiblePins = computed(() => allPins.value.slice(0, page.value * 12))

// ── Methods ──
function openPin(pin) {
  console.log('Open pin:', pin.id)
}

function savePin(pin) {
  console.log('Save pin:', pin.id)
}

function uploadPin() {
  console.log('Upload pin')
}

async function loadMore() {
  loading.value = true
  // Simulate API call
  await new Promise(resolve => setTimeout(resolve, 800))
  page.value++
  loading.value = false
}
</script>

<style scoped>
.home-page {
  min-height: 100vh;
  background: var(--bg);
  font-family: 'DM Sans', sans-serif;
  color: var(--ink);
}

/* ── Nav ── */


.avatar-nav {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: var(--ink);
  cursor: pointer;
  flex-shrink: 0;
  display: block;
}

/* ── Filters ── */
.filters-bar {
  background: var(--white);
  border-bottom: 1px solid var(--border);
  padding: 0 24px;
  padding-inline: 15%;
  display: flex;
  align-items: center;
  gap: 8px;
  overflow-x: auto;
  height: 52px;
  scrollbar-width: none;
}

.filters-bar::-webkit-scrollbar {
  display: none;
}

.filter-chip {
  padding: 6px 16px;
  border: 1.5px solid var(--border);
  border-radius: 100px;
  font-size: 13px;
  color: var(--muted);
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.2s;
  background: transparent;
  font-family: 'DM Sans', sans-serif;
}

.filter-chip:hover {
  border-color: var(--ink);
  color: var(--ink);
}

.filter-chip.active {
  background: var(--ink);
  color: var(--white);
  border-color: var(--ink);
}

/* ── Feed / Masonry ── */
.feed {
  padding: 24px;
  max-width: 1400px;
  margin: 0 auto;
}

.masonry {
  columns: 6;
  column-gap: 12px;
}

/* ── Load more ── */
.load-wrap {
  text-align: center;
  padding: 40px 0 60px;
}

.btn-load {
  padding: 12px 32px;
  border: 1.5px solid var(--ink);
  border-radius: 100px;
  background: transparent;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}

.btn-load:hover:not(:disabled) {
  background: var(--ink);
  color: var(--white);
}

.btn-load:disabled {
  opacity: 0.5;
  cursor: default;
}


@media (max-width: 1200px) {
  .masonry {
    columns: 5;
  }
}

@media (max-width: 960px) {
  .masonry {
    columns: 4;
  }
}

@media (max-width: 720px) {
  .masonry {
    columns: 3;
  }

  nav {
    padding: 0 16px;
    gap: 10px;
  }

  .nav-links {
    display: none;
  }

  .feed {
    padding: 16px;
  }
}

@media (max-width: 480px) {
  .masonry {
    columns: 2;
  }

  .btn-upload span {
    display: none;
  }
}
</style>
