<template>
  <div class="profile-page">
    <!-- ── Navigation ── -->

    <!-- ── Profile Header ── -->
    <header class="profile-header">
      <div class="avatar-wrap">
        <div class="avatar">{{ userInitials }}</div>
        <button class="avatar-edit" aria-label="Изменить фото">
          <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
          </svg>
        </button>
      </div>

      <h1 class="profile-name">{{ user.name }}</h1>
      <p class="profile-username">@{{ user.username }}</p>
      <p class="profile-bio">{{ user.bio }}</p>

      <div class="profile-stats">
        <div v-for="stat in stats" :key="stat.label" class="stat">
          <span class="stat-val">{{ stat.value }}</span>
          <span class="stat-label">{{ stat.label }}</span>
        </div>
      </div>

      <div class="profile-actions">
        <button class="btn-outline" @click="editProfile">Редактировать</button>
        <button class="btn-outline" @click="shareProfile">Поделиться</button>
      </div>

      <!-- Tabs -->
      <div class="tabs" role="tablist">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          class="tab"
          :class="{ active: activeTab === tab.id }"
          role="tab"
          :aria-selected="activeTab === tab.id"
          @click="activeTab = tab.id"
        >
          {{ tab.label }}
        </button>
      </div>
    </header>

    <!-- ── Content ── -->
    <main class="content">
      <!-- Boards tab -->
      <div v-if="activeTab === 'boards'" class="boards-grid">
        <div
          v-for="board in boards"
          :key="board.id"
          class="board-card"
          @click="openBoard(board)"
        >
          <div class="board-preview">
            <div :class="['board-img', board.colors[0]]" />
            <div :class="['board-img', board.colors[1]]" />
            <div :class="['board-img', board.colors[2]]" />
          </div>
          <div class="board-info">
            <span class="board-name">{{ board.name }}</span>
            <span class="board-count">{{ board.count }}</span>
          </div>
        </div>

        <button class="board-card board-new" @click="createBoard">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
            <circle cx="12" cy="12" r="10"/>
            <line x1="12" y1="8" x2="12" y2="16"/>
            <line x1="8" y1="12" x2="16" y2="12"/>
          </svg>
          <span>Новая доска</span>
        </button>
      </div>

      <!-- Pins tab -->
      <div v-else-if="activeTab === 'pins'" class="pins-placeholder">
        <p>Здесь будут ваши пины</p>
      </div>

      <!-- Likes tab -->
      <div v-else-if="activeTab === 'likes'" class="pins-placeholder">
        <p>Здесь будут ваши лайки</p>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { RouterLink } from 'vue-router'

// ── State ──
const searchQuery = ref('')
const activeTab = ref('boards')

const user = ref({
  name: 'Алексей Иванов',
  username: 'alexivanov',
  bio: 'Дизайнер интерьеров. Коллекционирую вдохновение из путешествий, архитектуры и современного искусства.',
  avatarUrl: 'https://api.dicebear.com/7.x/initials/svg?seed=AK&backgroundColor=1A1A18&textColor=FDFCFB',
})

const stats = [
  { label: 'пинов',       value: '248'  },
  { label: 'подписчиков', value: '1.2К' },
  { label: 'подписок',    value: '134'  },
]

const tabs = [
  { id: 'boards', label: 'Доски' },
  { id: 'pins',   label: 'Пины'  },
  { id: 'likes',  label: 'Лайки' },
]

const boards = ref([
  { id: 1, name: 'Интерьеры',   count: 64, colors: ['c1', 'c2', 'c3'] },
  { id: 2, name: 'Архитектура', count: 41, colors: ['c4', 'c5', 'c6'] },
  { id: 3, name: 'Минимализм',  count: 87, colors: ['c2', 'c1', 'c4'] },
  { id: 4, name: 'Цвет и форма',count: 33, colors: ['c5', 'c3', 'c6'] },
])

// ── Computed ──
const userInitials = computed(() => {
  return user.value.name
    .split(' ')
    .map(w => w[0])
    .join('')
    .slice(0, 2)
    .toUpperCase()
})

// ── Methods ──
function editProfile() {
  console.log('Edit profile')
}

function shareProfile() {
  navigator.clipboard?.writeText(window.location.href)
}

function openBoard(board) {
  console.log('Open board:', board.id)
}

function createBoard() {
  console.log('Create new board')
}
</script>

<style scoped>
.profile-page {
  min-height: 100vh;
  background: var(--bg);
  font-family: 'DM Sans', sans-serif;
  color: var(--ink);
}

/* ── Nav ── */

/* ── Profile Header ── */
.profile-header {
  background: var(--white);
  border-bottom: 1px solid var(--border);
  padding: 48px 40px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.avatar-wrap {
  position: relative;
  margin-bottom: 20px;
}

.avatar {
  width: 96px;
  height: 96px;
  border-radius: 50%;
  background: linear-gradient(135deg, #C8562A, #1A1A18);
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: 'DM Serif Display', serif;
  font-size: 36px;
  color: var(--white);
  letter-spacing: -1px;
}

.avatar-edit {
  position: absolute;
  bottom: 0;
  right: 0;
  width: 28px;
  height: 28px;
  background: var(--white);
  border: 1.5px solid var(--border);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: var(--muted);
  transition: border-color 0.2s, color 0.2s;
}

.avatar-edit:hover { border-color: var(--ink); color: var(--ink); }

.profile-name {
  font-family: 'DM Serif Display', serif;
  font-size: 28px;
  letter-spacing: -0.6px;
  margin-bottom: 4px;
}

.profile-username {
  font-size: 14px;
  color: var(--muted);
  margin-bottom: 12px;
}

.profile-bio {
  font-size: 14px;
  color: var(--muted);
  max-width: 380px;
  line-height: 1.6;
  font-weight: 300;
  margin-bottom: 20px;
}

.profile-stats {
  display: flex;
  gap: 36px;
  margin-bottom: 24px;
}

.stat {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
}

.stat-val {
  font-family: 'DM Serif Display', serif;
  font-size: 20px;
  letter-spacing: -0.4px;
}

.stat-label {
  font-size: 12px;
  color: var(--muted);
  font-weight: 300;
}

.profile-actions {
  display: flex;
  gap: 10px;
  margin-bottom: 28px;
}

.btn-outline {
  padding: 9px 22px;
  border: 1.5px solid var(--border);
  border-radius: 8px;
  background: transparent;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  color: var(--ink);
  cursor: pointer;
  transition: border-color 0.2s;
}

.btn-outline:hover { border-color: var(--ink); }

/* ── Tabs ── */
.tabs {
  display: flex;
  border-top: 1px solid var(--border);
  width: 100%;
  justify-content: center;
}

.tab {
  padding: 14px 24px;
  font-size: 13px;
  font-weight: 500;
  color: var(--muted);
  border: none;
  border-bottom: 2px solid transparent;
  background: transparent;
  cursor: pointer;
  letter-spacing: 0.02em;
  transition: color 0.2s;
  margin-bottom: -1px;
  font-family: 'DM Sans', sans-serif;
}

.tab.active {
  color: var(--ink);
  border-bottom-color: var(--ink);
}

.tab:hover:not(.active) { color: var(--ink); }

/* ── Content ── */
.content {
  padding: 32px 40px;
  max-width: 1200px;
  margin: 0 auto;
}

.boards-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 16px;
}

.board-card {
  background: var(--white);
  border: 1px solid var(--border);
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  transition: border-color 0.2s, transform 0.2s;
}

.board-card:hover { border-color: var(--ink); transform: translateY(-2px); }

.board-preview {
  display: grid;
  grid-template-columns: 2fr 1fr;
  grid-template-rows: 1fr 1fr;
  height: 130px;
  gap: 2px;
}

.board-img { background: var(--border); }
.board-img:first-child { grid-row: span 2; }

.c1 { background: #D4CAC0; }
.c2 { background: #E8D5C4; }
.c3 { background: #C8B8A8; }
.c4 { background: #DDD0C0; }
.c5 { background: #B8C8C0; }
.c6 { background: #C8D4D8; }

.board-info {
  padding: 12px 14px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.board-name {
  font-size: 14px;
  font-weight: 500;
}

.board-count {
  font-size: 12px;
  color: var(--muted);
}

.board-new {
  border: 2px dashed var(--border);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 8px;
  color: var(--muted);
  font-size: 13px;
  min-height: 178px;
  background: transparent;
  transition: border-color 0.2s, color 0.2s;
  border-radius: 12px;
}

.board-new:hover { border-color: var(--ink); color: var(--ink); }

.pins-placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 200px;
  color: var(--muted);
  font-size: 14px;
}

@media (max-width: 768px) {
  nav { padding: 0 20px; }
  .nav-search { max-width: 200px; margin: 0 16px; }
  .profile-header { padding: 32px 20px 0; }
  .content { padding: 24px 20px; }
  .boards-grid { grid-template-columns: repeat(2, 1fr); }
}

@media (max-width: 480px) {
  .profile-stats { gap: 24px; }
}
</style>
