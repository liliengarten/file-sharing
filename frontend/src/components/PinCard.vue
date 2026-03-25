<script setup>
const props = defineProps({
  pin: {
    type: Object,
    required: true
  }
})
</script>

<template>
  <article
      class="pin"
      :style="{ animationDelay: `${i * 0.04}s` }">
    <div class="pin-img-wrap">
      <div :class="['pin-img', pin.height, pin.palette]" />
      <div class="pin-overlay">
        <button class="pin-save" @click.stop="savePin(pin)">
          <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/>
          </svg>
          Сохранить
        </button>
      </div>
    </div>
    <div class="pin-body">
      <p class="pin-title">{{ pin.title }}</p>
      <div class="pin-meta">
        <div class="pin-author">
          <div class="author-dot">{{ pin.authorInitials }}</div>
          {{ pin.author }}
        </div>
        <div class="pin-likes">
          <svg width="11" height="11" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
          </svg>
          {{ pin.likes }}
        </div>
      </div>
    </div>
  </article>
</template>

<style scoped>
.pin {
  break-inside: avoid;
  background: var(--white);
  border-radius: 12px;
  overflow: hidden;
  margin-bottom: 12px;
  cursor: pointer;
  border: 1px solid transparent;
  position: relative;
  transition: transform 0.25s cubic-bezier(0.34, 1.56, 0.64, 1);
  animation: pinIn 0.4s ease both;
}

@keyframes pinIn {
  from { opacity: 0; transform: translateY(12px); }
  to   { opacity: 1; transform: translateY(0); }
}

.pin:hover { transform: translateY(-4px); border-color: var(--border); }
.pin:hover .pin-overlay { opacity: 1; }
.pin:hover .pin-img { transform: scale(1.03); }

.pin-img-wrap {
  position: relative;
  overflow: hidden;
}

.pin-img {
  width: 100%;
  display: block;
  transition: transform 0.3s;
}

/* Heights */
.h-xs  { height: 120px; }
.h-sm  { height: 160px; }
.h-md  { height: 210px; }
.h-lg  { height: 260px; }
.h-xl  { height: 320px; }

/* Palettes */
.pal-1  { background: #D4C8B8; }
.pal-2  { background: #B8C8C0; }
.pal-3  { background: #C8B8A0; }
.pal-4  { background: #D8D0C8; }
.pal-5  { background: #C0C8D8; }
.pal-6  { background: #D0C0C8; }
.pal-7  { background: #C8D0C0; }
.pal-8  { background: #D0C8B0; }
.pal-9  { background: #B8C0C8; }
.pal-10 { background: #C8C0B8; }

.pin-overlay {
  position: absolute;
  inset: 0;
  background: rgba(26, 26, 24, 0.08);
  opacity: 0;
  transition: opacity 0.2s;
  display: flex;
  align-items: flex-start;
  justify-content: flex-end;
  padding: 10px;
}

.pin-save {
  background: var(--white);
  border: none;
  border-radius: 100px;
  padding: 7px 14px;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 5px;
  transition: background 0.2s;
  color: var(--ink);
}

.pin-save:hover { background: var(--bg); }

.pin-body { padding: 10px 12px 12px; }

.pin-title {
  font-size: 13px;
  font-weight: 500;
  line-height: 1.4;
  color: var(--ink);
  margin-bottom: 6px;
}

.pin-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.pin-author {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  color: var(--muted);
}

.author-dot {
  width: 18px;
  height: 18px;
  border-radius: 50%;
  background: var(--border);
  font-size: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--muted);
  font-weight: 500;
}

.pin-likes {
  display: flex;
  align-items: center;
  gap: 3px;
  font-size: 11px;
  color: var(--muted);
}
</style>