<template>
  <div class="register-layout">
    <!-- Left visual panel -->
    <div class="panel-visual">
      <div class="visual-grid">
        <div v-for="n in 10" :key="n" class="visual-card" />
      </div>
      <div class="logo">pin<span>flo</span></div>
      <div class="tagline">
        <h2>Делитесь тем,<br>что вдохновляет</h2>
        <p>Коллекционируйте изображения, создавайте доски и находите своё сообщество.</p>
      </div>
    </div>

    <!-- Right form panel -->
    <div class="panel-form">
      <div class="form-wrap">
        <h1 class="form-title">Создать аккаунт</h1>
        <p class="form-sub">
          Уже есть аккаунт?
          <routerLink to="/login">Войти</routerLink>
        </p>

        <form @submit.prevent="handleSubmit">
          <div class="row-2">
            <div class="field">
              <label for="firstName">Имя</label>
              <input
                id="firstName"
                v-model="form.firstName"
                type="text"
                placeholder="Алексей"
              />
            </div>
            <div class="field">
              <label for="lastName">Фамилия</label>
              <input
                id="lastName"
                v-model="form.lastName"
                type="text"
                placeholder="Иванов"
              />
            </div>
          </div>

          <div class="field">
            <label for="username">Имя пользователя</label>
            <input
              id="username"
              v-model="form.username"
              type="text"
              placeholder="@alexivanov"
            />
          </div>

          <div class="field">
            <label for="email">Email</label>
            <input
              id="email"
              v-model="form.email"
              type="email"
              placeholder="alex@email.com"
            />
          </div>

          <div class="field">
            <label for="password">Пароль</label>
            <input
              id="password"
              v-model="form.password"
              type="password"
              placeholder="Минимум 8 символов"
            />
          </div>

          <button type="submit" class="btn-primary">Создать аккаунт</button>
        </form>

        <div class="divider">или</div>

        <p class="terms">
          Регистрируясь, вы принимаете
          <a href="#">Условия использования</a><br>
          и <a href="#">Политику конфиденциальности</a>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { RouterLink } from 'vue-router'

const form = reactive({
  firstName: '',
  lastName: '',
  username: '',
  email: '',
  password: '',
})

function handleSubmit() {
  // TODO: connect to your auth API
  console.log('Register:', form)
}

function handleGoogle() {
  // TODO: OAuth flow
  console.log('Google sign-up')
}

// Inline SVG as component for cleanliness
function GoogleIcon() {}
</script>

<script>
// GoogleIcon as a functional sub-component
import { defineComponent, h } from 'vue'

export const GoogleIcon = defineComponent({
  name: 'GoogleIcon',
  render() {
    return h('svg', { width: 18, height: 18, viewBox: '0 0 24 24', fill: 'none' }, [
      h('path', { d: 'M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z', fill: '#4285F4' }),
      h('path', { d: 'M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z', fill: '#34A853' }),
      h('path', { d: 'M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z', fill: '#FBBC05' }),
      h('path', { d: 'M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z', fill: '#EA4335' }),
    ])
  },
})
</script>

<style scoped>
.register-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  min-height: 100vh;
  font-family: 'DM Sans', sans-serif;
  color: var(--ink);
}

/* ── LEFT PANEL ── */
.panel-visual {
  background: var(--ink);
  position: relative;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 7%;
}

.panel-visual::before {
  content: '';
  position: absolute;
  width: 480px;
  height: 480px;
  border-radius: 50%;
  background: radial-gradient(circle, #C8562A22 0%, transparent 70%);
  top: -120px;
  left: -120px;
}

.visual-grid {
  position: absolute;
  inset: 0;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(4, 1fr);
  gap: 8px;
  padding: 48px;
  opacity: 0.12;
  animation: fadeIn 1s ease forwards;
}

.visual-card {
  background: #fff;
  border-radius: 8px;
}

.visual-card:nth-child(2) { grid-row: span 2; background: #C8562A; opacity: 0.6; }
.visual-card:nth-child(5) { grid-row: span 2; }
.visual-card:nth-child(8) { grid-column: span 2; }

.logo {
  font-family: 'DM Serif Display', serif;
  font-size: 28px;
  color: var(--white);
  letter-spacing: -0.5px;
  position: relative;
  z-index: 2;
}

.logo span { color: var(--accent); }

.tagline {
  position: relative;
  z-index: 2;
}

.tagline h2 {
  font-family: 'DM Serif Display', serif;
  font-size: 42px;
  line-height: 1.15;
  color: var(--white);
  margin-bottom: 16px;
  letter-spacing: -1px;
}

.tagline p {
  font-size: 15px;
  color: #8A8880;
  font-weight: 300;
  line-height: 1.6;
  max-width: 300px;
}

/* ── RIGHT PANEL ── */
.panel-form {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 64px 56px;
  background: var(--white);
}

.form-wrap {
  width: 100%;
  max-width: 360px;
  animation: slideUp 0.6s ease forwards;
}

.form-title {
  font-family: 'DM Serif Display', serif;
  font-size: 34px;
  letter-spacing: -0.8px;
  margin-bottom: 6px;
  color: var(--ink);
}

.form-sub {
  font-size: 14px;
  color: var(--muted);
  margin-bottom: 36px;
  font-weight: 300;
}

.form-sub a {
  color: var(--accent);
  text-decoration: none;
  font-weight: 500;
}

.field { margin-bottom: 18px; }

label {
  display: block;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  color: var(--muted);
  margin-bottom: 8px;
}

input {
  width: 100%;
  padding: 13px 16px;
  border: 1.5px solid var(--border);
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 15px;
  background: var(--bg);
  color: var(--ink);
  outline: none;
  transition: border-color 0.2s, background 0.2s;
  box-sizing: border-box;
}

input::placeholder { color: #C4C0BA; }

input:focus {
  border-color: var(--ink);
  background: var(--white);
}

.row-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.btn-primary {
  width: 100%;
  padding: 14px;
  background: var(--ink);
  color: var(--white);
  border: none;
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  margin-top: 8px;
  transition: background 0.2s, transform 0.1s;
}

.btn-primary:hover { background: #2c2c2a; }
.btn-primary:active { transform: scale(0.99); }

.divider {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 24px 0;
  font-size: 12px;
  color: var(--muted);
}

.divider::before,
.divider::after {
  content: '';
  flex: 1;
  height: 1px;
  background: var(--border);
}

.btn-social {
  width: 100%;
  padding: 13px;
  background: transparent;
  color: var(--ink);
  border: 1.5px solid var(--border);
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  font-weight: 400;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  transition: border-color 0.2s;
}

.btn-social:hover { border-color: var(--ink); }

.terms {
  font-size: 11px;
  color: var(--muted);
  text-align: center;
  margin-top: 20px;
  line-height: 1.7;
}

.terms a { color: var(--ink); text-decoration: none; }

@keyframes fadeIn {
  from { opacity: 0; }
  to   { opacity: 0.12; }
}

@keyframes slideUp {
  from { opacity: 0; transform: translateY(20px); }
  to   { opacity: 1; transform: translateY(0); }
}

@media (max-width: 768px) {
  .register-layout { grid-template-columns: 1fr; }
  .panel-visual { display: none; }
  .panel-form { padding: 40px 24px; }
}
</style>
