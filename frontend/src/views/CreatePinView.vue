<template>
  <div class="create-page">

    <!-- ── Editor ── -->
    <div class="editor">

      <!-- ── Upload zone ── -->
      <div
        class="upload-zone"
        :class="{ 'has-file': previewUrl, 'drag-over': isDragging }"
        @dragover.prevent="isDragging = true"
        @dragleave="isDragging = false"
        @drop.prevent="onDrop"
        @click="!previewUrl && triggerFileInput()"
      >
        <!-- Empty state -->
        <template v-if="!previewUrl">
          <div class="upload-icon">
            <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
              <polyline points="17 8 12 3 7 8"/>
              <line x1="12" y1="3" x2="12" y2="15"/>
            </svg>
          </div>
          <p class="upload-label">Перетащите файл или нажмите, чтобы выбрать</p>
          <p class="upload-hint">PNG, JPG, GIF, WEBP · до 20 МБ</p>
          <button class="btn-choose" @click.stop="triggerFileInput">Выбрать файл</button>
        </template>

        <!-- Preview -->
        <template v-else>
          <img :src="previewUrl" alt="Preview" class="preview-img" />
          <div class="preview-overlay">
            <button class="btn-replace" @click.stop="triggerFileInput">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                <polyline points="17 8 12 3 7 8"/>
                <line x1="12" y1="3" x2="12" y2="15"/>
              </svg>
              Заменить
            </button>
            <button class="btn-remove" @click.stop="removeFile">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polyline points="3 6 5 6 21 6"/>
                <path d="M19 6l-1 14H6L5 6"/>
                <path d="M10 11v6"/><path d="M14 11v6"/>
                <path d="M9 6V4h6v2"/>
              </svg>
              Удалить
            </button>
          </div>
        </template>

        <input
          ref="fileInputRef"
          type="file"
          accept="image/*,video/*"
          class="file-input"
          @change="onFileChange"
        />
      </div>

      <!-- ── Form ── -->
      <div class="form-side">

        <!-- Title -->
        <div class="field">
          <div class="field-header">
            <label>Название</label>
            <span class="char-count">{{ form.title.length }} / 100</span>
          </div>
          <input
            v-model="form.title"
            type="text"
            class="input"
            placeholder="Добавьте заголовок..."
            maxlength="100"
          />
        </div>

        <!-- Description -->
        <div class="field">
          <div class="field-header">
            <label>Описание</label>
            <span class="char-count">{{ form.description.length }} / 500</span>
          </div>
          <textarea
            v-model="form.description"
            class="textarea"
            placeholder="Расскажите подробнее..."
            maxlength="500"
            rows="4"
          />
        </div>

        <!-- Source URL -->
        <div class="field">
          <label>Ссылка на источник</label>
          <div class="input-icon-wrap">
            <svg class="input-icon" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/>
              <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>
            </svg>
            <input
              v-model="form.sourceUrl"
              type="url"
              class="input with-icon"
              placeholder="https://example.com"
            />
          </div>
        </div>

        <!-- Board -->
        <div class="field">
          <label>Доска</label>
          <div class="select-wrap">
            <select v-model="form.board" class="select">
              <option value="">Выберите доску...</option>
              <option v-for="board in boards" :key="board.id" :value="board.id">
                {{ board.name }}
              </option>
            </select>
            <svg class="select-arrow" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="6 9 12 15 18 9"/>
            </svg>
          </div>
          <button class="btn-new-board" @click="showBoardModal = true">
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="16"/><line x1="8" y1="12" x2="16" y2="12"/>
            </svg>
            Создать доску
          </button>
        </div>

        <!-- Tags -->
        <div class="field">
          <label>Теги</label>
          <div class="tags-input-wrap" @click="tagInputRef?.focus()">
            <span v-for="tag in form.tags" :key="tag" class="tag">
              #{{ tag }}
              <button class="tag-remove" @click.stop="removeTag(tag)">×</button>
            </span>
            <input
              ref="tagInputRef"
              v-model="tagInput"
              type="text"
              class="tag-input"
              placeholder="Добавить тег..."
              @keydown.enter.prevent="addTag"
              @keydown.comma.prevent="addTag"
              @keydown.backspace="onTagBackspace"
            />
          </div>
          <p class="field-hint">Нажмите Enter или запятую для добавления</p>
        </div>

        <!-- Alt text -->
        <div class="field">
          <div class="field-header">
            <label>Альтернативный текст</label>
            <span class="badge-optional">Необязательно</span>
          </div>
          <input
            v-model="form.altText"
            type="text"
            class="input"
            placeholder="Опишите изображение для людей с нарушениями зрения"
          />
        </div>

        <!-- Visibility -->
        <div class="field">
          <label>Видимость</label>
          <div class="toggle-group">
            <button
              v-for="opt in visibilityOptions"
              :key="opt.value"
              class="toggle-btn"
              :class="{ active: form.visibility === opt.value }"
              @click="form.visibility = opt.value"
            >
              <component :is="opt.icon" />
              {{ opt.label }}
            </button>
          </div>
        </div>

        <!-- Publish date -->
        <div class="field">
          <div class="field-header">
            <label>Дата публикации</label>
            <span class="badge-optional">Необязательно</span>
          </div>
          <div class="radio-group">
            <label class="radio-label">
              <input v-model="form.publishType" type="radio" value="now" />
              <span>Сейчас</span>
            </label>
            <label class="radio-label">
              <input v-model="form.publishType" type="radio" value="later" />
              <span>Запланировать</span>
            </label>

          </div>
          <input
            v-if="form.publishType === 'later'"
            v-model="form.publishDate"
            type="datetime-local"
            class="input"
            style="margin-top: 10px;"
          />
        </div>
        <div class="nav-right">
          <button class="btn-draft" @click="saveDraft">Сохранить черновик</button>
          <button
              class="btn-publish"
              :disabled="!canPublish"
              @click="publish"
          >
            Опубликовать
          </button>
        </div>

        <!-- Bottom actions (mobile) -->
        <div class="form-actions-mobile">
          <button class="btn-draft-mobile" @click="saveDraft">Черновик</button>
          <button
            class="btn-publish-mobile"
            :disabled="!canPublish"
            @click="publish"
          >
            Опубликовать
          </button>
        </div>

      </div>
    </div>

    <!-- ── New Board Modal ── -->
    <Transition name="modal">
      <div v-if="showBoardModal" class="modal-overlay" @click.self="showBoardModal = false">
        <div class="modal">
          <div class="modal-header">
            <h2>Новая доска</h2>
            <button class="modal-close" @click="showBoardModal = false">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>
          <div class="field">
            <label>Название доски</label>
            <input v-model="newBoardName" type="text" class="input" placeholder="Интерьеры, путешествия..." />
          </div>
          <div class="field">
            <label>Видимость</label>
            <div class="toggle-group">
              <button
                class="toggle-btn"
                :class="{ active: newBoardPrivate === false }"
                @click="newBoardPrivate = false"
              >Публичная</button>
              <button
                class="toggle-btn"
                :class="{ active: newBoardPrivate === true }"
                @click="newBoardPrivate = true"
              >Приватная</button>
            </div>
          </div>
          <div class="modal-footer">
            <button class="btn-cancel" @click="showBoardModal = false">Отмена</button>
            <button class="btn-create-board" :disabled="!newBoardName.trim()" @click="createBoard">Создать</button>
          </div>
        </div>
      </div>
    </Transition>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { RouterLink, useRouter } from 'vue-router'

const router = useRouter()

// ── Refs ──
const fileInputRef = ref(null)
const tagInputRef  = ref(null)

// ── Upload state ──
const previewUrl  = ref('')
const isDragging  = ref(false)
const selectedFile = ref(null)

// ── Form ──
const form = ref({
  title:       '',
  description: '',
  sourceUrl:   '',
  board:       '',
  tags:        [],
  altText:     '',
  visibility:  'public',
  publishType: 'now',
  publishDate: '',
})

const tagInput = ref('')

// ── Boards ──
const boards = ref([
  { id: 1, name: 'Интерьеры' },
  { id: 2, name: 'Архитектура' },
  { id: 3, name: 'Минимализм' },
  { id: 4, name: 'Цвет и форма' },
])

// ── Modal ──
const showBoardModal  = ref(false)
const newBoardName    = ref('')
const newBoardPrivate = ref(false)

// ── Computed ──
const canPublish = computed(() =>
  !!previewUrl.value && form.value.title.trim().length > 0
)

const visibilityOptions = [
  { value: 'public',  label: 'Для всех'  },
  { value: 'private', label: 'Только я'  },
]

// ── File handling ──
function triggerFileInput() {
  fileInputRef.value?.click()
}

function onFileChange(e) {
  const file = e.target.files?.[0]
  if (file) processFile(file)
}

function onDrop(e) {
  isDragging.value = false
  const file = e.dataTransfer.files?.[0]
  if (file) processFile(file)
}

function processFile(file) {
  selectedFile.value = file
  const reader = new FileReader()
  reader.onload = (e) => { previewUrl.value = e.target.result }
  reader.readAsDataURL(file)
}

function removeFile() {
  previewUrl.value = ''
  selectedFile.value = null
  if (fileInputRef.value) fileInputRef.value.value = ''
}

// ── Tags ──
function addTag() {
  const tag = tagInput.value.trim().replace(/^#/, '').replace(/,/g, '')
  if (tag && !form.value.tags.includes(tag) && form.value.tags.length < 10) {
    form.value.tags.push(tag)
  }
  tagInput.value = ''
}

function removeTag(tag) {
  form.value.tags = form.value.tags.filter(t => t !== tag)
}

function onTagBackspace() {
  if (!tagInput.value && form.value.tags.length) {
    form.value.tags.pop()
  }
}

// ── Board ──
function createBoard() {
  if (!newBoardName.value.trim()) return
  const id = Date.now()
  boards.value.push({ id, name: newBoardName.value.trim() })
  form.value.board = id
  newBoardName.value = ''
  showBoardModal.value = false
}

// ── Actions ──
function saveDraft() {
  console.log('Draft saved', form.value)
  // In real app: call API / store in localStorage
}

function publish() {
  if (!canPublish.value) return
  console.log('Publishing pin', { file: selectedFile.value, ...form.value })
  // In real app: POST to API, then navigate
  router.push('/profile')
}
</script>

<style scoped>
.create-page {
  min-height: 100vh;
  background: var(--bg);
  font-family: 'DM Sans', sans-serif;
  color: var(--ink);
}

/* ── Nav ── */
nav {
  position: sticky;
  top: 0;
  background: var(--white);
  border-bottom: 1px solid var(--border);
  z-index: 100;
  padding: 0 24px;
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

.nav-center {
  flex: 1;
  display: flex;
  justify-content: center;
}

.nav-title {
  font-size: 15px;
  font-weight: 500;
  letter-spacing: -0.2px;
}

.nav-right {
  display: flex;
  align-items: center;
  gap: 10px;
  flex-shrink: 0;
}

.btn-draft {
  padding: 8px 18px;
  background: transparent;
  color: var(--ink);
  border: 1.5px solid var(--border);
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: border-color 0.2s;
  white-space: nowrap;
}
.btn-draft:hover { border-color: var(--ink); }

.btn-publish {
  padding: 8px 22px;
  background: var(--accent);
  color: var(--white);
  border: none;
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: opacity 0.2s;
  white-space: nowrap;
}
.btn-publish:hover:not(:disabled) { opacity: 0.85; }
.btn-publish:disabled { opacity: 0.35; cursor: not-allowed; }

/* ── Editor layout ── */
.editor {
  display: grid;
  grid-template-columns: 420px 1fr;
  gap: 0;
  max-width: 1100px;
  margin: 48px auto;
  background: var(--white);
  border: 1px solid var(--border);
  border-radius: 20px;
  overflow: hidden;
  animation: editorIn 0.4s ease both;
}

/* ── Upload zone ── */
.upload-zone {
  position: relative;
  background: var(--bg);
  min-height: 600px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 12px;
  cursor: pointer;
  border-right: 1px solid var(--border);
  transition: background 0.2s;
  overflow: hidden;
}

.upload-zone.drag-over { background: #EDE8E0; }
.upload-zone.has-file  { cursor: default; }

.upload-icon {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  background: var(--white);
  border: 1.5px solid var(--border);
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--muted);
  margin-bottom: 4px;
}

.upload-label {
  font-size: 15px;
  font-weight: 500;
  color: var(--ink);
  text-align: center;
  padding: 0 32px;
}

.upload-hint {
  font-size: 12px;
  color: var(--muted);
  font-weight: 300;
}

.btn-choose {
  margin-top: 8px;
  padding: 9px 22px;
  background: var(--ink);
  color: var(--white);
  border: none;
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
}
.btn-choose:hover { background: #2c2c2a; }

.file-input {
  display: none;
}

/* Preview */
.preview-img {
  width: 100%;
  height: 100%;
  object-fit: contain;
  position: absolute;
  inset: 0;
}

.preview-overlay {
  position: absolute;
  bottom: 16px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 8px;
  z-index: 2;
  opacity: 0;
  transition: opacity 0.2s;
}

.upload-zone:hover .preview-overlay { opacity: 1; }

.btn-replace,
.btn-remove {
  padding: 8px 14px;
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
  transition: opacity 0.2s;
}

.btn-replace {
  background: var(--white);
  color: var(--ink);
  border: 1.5px solid var(--border);
}
.btn-replace:hover { border-color: var(--ink); }

.btn-remove {
  background: #fff0ed;
  color: var(--accent);
  border: 1.5px solid #f4c4b4;
}
.btn-remove:hover { opacity: 0.8; }

/* ── Form side ── */
.form-side {
  padding: 36px 40px;
  display: flex;
  flex-direction: column;
  gap: 24px;
  overflow-y: auto;
  max-height: 90vh;
}

.field { display: flex; flex-direction: column; gap: 8px; }

.field-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

label {
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  color: var(--muted);
}

.char-count {
  font-size: 11px;
  color: var(--muted);
  font-weight: 300;
}

.badge-optional {
  font-size: 11px;
  color: var(--muted);
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: 4px;
  padding: 2px 6px;
  font-weight: 400;
  text-transform: none;
  letter-spacing: 0;
}

.field-hint {
  font-size: 11px;
  color: var(--muted);
  font-weight: 300;
}

.input,
.textarea,
.select {
  width: 100%;
  padding: 11px 14px;
  border: 1.5px solid var(--border);
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  background: var(--bg);
  color: var(--ink);
  outline: none;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.input:focus,
.textarea:focus,
.select:focus { border-color: var(--ink); background: var(--white); }

.input::placeholder,
.textarea::placeholder { color: #C4C0BA; }

.textarea { resize: vertical; min-height: 100px; }

/* Input with icon */
.input-icon-wrap { position: relative; }
.input-icon {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--muted);
  pointer-events: none;
}
.input.with-icon { padding-left: 34px; }

/* Select */
.select-wrap { position: relative; }

.select {
  appearance: none;
  padding-right: 32px;
  cursor: pointer;
}

.select-arrow {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  color: var(--muted);
  pointer-events: none;
}

.btn-new-board {
  align-self: flex-start;
  display: flex;
  align-items: center;
  gap: 5px;
  font-family: 'DM Sans', sans-serif;
  font-size: 12px;
  font-weight: 500;
  color: var(--muted);
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  transition: color 0.2s;
}
.btn-new-board:hover { color: var(--ink); }

/* Tags input */
.tags-input-wrap {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  padding: 10px 12px;
  border: 1.5px solid var(--border);
  border-radius: 8px;
  background: var(--bg);
  cursor: text;
  min-height: 46px;
  align-items: center;
  transition: border-color 0.2s;
}

.tags-input-wrap:focus-within { border-color: var(--ink); background: var(--white); }

.tag {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 3px 8px;
  background: var(--white);
  border: 1px solid var(--border);
  border-radius: 100px;
  font-size: 12px;
  color: var(--ink);
}

.tag-remove {
  background: none;
  border: none;
  cursor: pointer;
  color: var(--muted);
  font-size: 14px;
  line-height: 1;
  padding: 0;
  display: flex;
  align-items: center;
}
.tag-remove:hover { color: var(--accent); }

.tag-input {
  border: none;
  outline: none;
  background: transparent;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  color: var(--ink);
  min-width: 120px;
  flex: 1;
}
.tag-input::placeholder { color: #C4C0BA; }

/* Toggle group */
.toggle-group {
  display: flex;
  gap: 6px;
}

.toggle-btn {
  padding: 8px 18px;
  border: 1.5px solid var(--border);
  border-radius: 8px;
  background: transparent;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  color: var(--muted);
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  gap: 6px;
}
.toggle-btn:hover { border-color: var(--ink); color: var(--ink); }
.toggle-btn.active { background: var(--ink); color: var(--white); border-color: var(--ink); }

/* Radio */
.radio-group { display: flex; gap: 20px; }

.radio-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: var(--ink);
  cursor: pointer;
  text-transform: none;
  letter-spacing: 0;
  font-weight: 400;
}

.radio-label input[type="radio"] {
  accent-color: var(--ink);
  width: 15px;
  height: 15px;
  cursor: pointer;
}

/* Mobile actions */
.form-actions-mobile { display: none; }

/* ── Modal ── */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(26, 26, 24, 0.4);
  z-index: 200;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}

.modal {
  background: var(--white);
  border-radius: 16px;
  padding: 32px;
  width: 100%;
  max-width: 380px;
  display: flex;
  flex-direction: column;
  gap: 20px;
  border: 1px solid var(--border);
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.modal-header h2 {
  font-family: 'DM Serif Display', serif;
  font-size: 22px;
  letter-spacing: -0.4px;
}

.modal-close {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  border: 1.5px solid var(--border);
  background: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--muted);
  transition: border-color 0.2s, color 0.2s;
}
.modal-close:hover { border-color: var(--ink); color: var(--ink); }

.modal-footer {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}

.btn-cancel {
  padding: 9px 18px;
  border: 1.5px solid var(--border);
  border-radius: 8px;
  background: transparent;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  cursor: pointer;
  transition: border-color 0.2s;
}
.btn-cancel:hover { border-color: var(--ink); }

.btn-create-board {
  padding: 9px 22px;
  background: var(--ink);
  color: var(--white);
  border: none;
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
}
.btn-create-board:hover:not(:disabled) { background: #2c2c2a; }
.btn-create-board:disabled { opacity: 0.35; cursor: not-allowed; }

/* ── Transitions ── */
.modal-enter-active,
.modal-leave-active { transition: opacity 0.2s, transform 0.2s; }
.modal-enter-from,
.modal-leave-to { opacity: 0; transform: scale(0.96); }

@keyframes editorIn {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}

/* ── Responsive ── */
@media (max-width: 860px) {
  .editor {
    grid-template-columns: 1fr;
    margin: 0;
    border-radius: 0;
    border-left: none;
    border-right: none;
    border-top: none;
  }

  .upload-zone {
    min-height: 280px;
    border-right: none;
    border-bottom: 1px solid var(--border);
  }

  .form-side { max-height: none; padding: 28px 20px; }
  .form-actions-mobile { display: flex; gap: 10px; }

  .btn-draft-mobile,
  .btn-publish-mobile {
    flex: 1;
    padding: 13px;
    border-radius: 8px;
    font-family: 'DM Sans', sans-serif;
    font-size: 14px;
    font-weight: 500;
    cursor: pointer;
  }

  .btn-draft-mobile {
    background: transparent;
    border: 1.5px solid var(--border);
    color: var(--ink);
  }

  .btn-publish-mobile {
    background: var(--accent);
    color: var(--white);
    border: none;
    transition: opacity 0.2s;
  }
  .btn-publish-mobile:disabled { opacity: 0.35; cursor: not-allowed; }

  .nav-right { display: none; }
}
</style>
