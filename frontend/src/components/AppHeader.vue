<script setup>

import {RouterLink} from "vue-router";
</script>

<template>
  <nav>
    <RouterLink to="/" class="logo">pin<span>flo</span></RouterLink>

    <div class="nav-links">
      <RouterLink
          v-for="link in navLinks"
          :key="link.to"
          :to="link.to"
          class="nav-link"
          :class="{ active: link.to === $route.path }"
      >
        {{ link.label }}
      </RouterLink>
    </div>

    <div class="nav-search">
      <svg class="search-icon" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <circle cx="11" cy="11" r="8"/><path d="m21 21-4.35-4.35"/>
      </svg>
      <input v-model="searchQuery" type="text" placeholder="Найти идеи, авторов, темы..." />
    </div>

    <div class="nav-right">
      <button class="btn-upload" @click="uploadPin">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
          <line x1="12" y1="5" x2="12" y2="19"/>
          <line x1="5" y1="12" x2="19" y2="12"/>
        </svg>
        <span>Загрузить</span>
      </button>
      <button class="btn-icon" aria-label="Уведомления">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"/>
          <path d="M13.73 21a2 2 0 0 1-3.46 0"/>
        </svg>
      </button>
      <router-link to="/profile">Профиль</router-link>
    </div>
  </nav>
</template>

<style scoped>
nav {
  position: sticky;
  top: 0;
  background: var(--white);
  border-bottom: 1px solid var(--border);
  z-index: 100;
  padding: 0 24px;
  padding-inline: 15%;
  display: flex;
  align-items: center;
  gap: 16px;
  height: 60px;
}

.logo {
  font-family: 'DM Serif Display', serif;
  font-size: 22px;
  color: var(--ink);
  letter-spacing: -0.5px;
  text-decoration: none;
  flex-shrink: 0;
}

.logo span { color: var(--accent); }

.nav-links {
  display: flex;
  gap: 4px;
  flex-shrink: 0;
}

.nav-link {
  padding: 7px 14px;
  border-radius: 100px;
  font-size: 14px;
  font-weight: 500;
  color: var(--muted);
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
  text-decoration: none;
  white-space: nowrap;
}

.nav-link:hover { background: var(--bg); color: var(--ink); }
.nav-link.active { background: var(--ink); color: var(--white); }

.nav-search {
  flex: 1;
  position: relative;
  max-width: 480px;
}

.nav-search input {
  width: 100%;
  padding: 9px 16px 9px 38px;
  border: 1.5px solid var(--border);
  border-radius: 100px;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  background: var(--bg);
  color: var(--ink);
  outline: none;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.nav-search input:focus { border-color: var(--ink); }
.nav-search input::placeholder { color: #C4C0BA; }

.search-icon {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--muted);
  pointer-events: none;
}

.nav-right {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-left: auto;
  flex-shrink: 0;
}

.btn-upload {
  padding: 8px 18px;
  background: var(--accent);
  color: var(--white);
  border: none;
  border-radius: 100px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: opacity 0.2s;
  white-space: nowrap;
}

.btn-upload:hover { opacity: 0.85; }

.btn-icon {
  width: 38px;
  height: 38px;
  border: 1.5px solid var(--border);
  border-radius: 50%;
  background: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--ink);
  transition: border-color 0.2s;
  flex-shrink: 0;
}

.btn-icon:hover { border-color: var(--ink); }
</style>