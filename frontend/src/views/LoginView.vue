<template>
  <div class="login-page">
    <!-- Background decorations -->
    <div class="bg-decor" aria-hidden="true">
      <div class="circle-1" />
      <div class="circle-2" />
    </div>

    <div class="card">
      <RouterLink to="/" class="logo">pin<span>flo</span></RouterLink>

      <h1>С возвращением</h1>
      <p class="sub">
        Нет аккаунта?
        <routerLink to="/register">Создать</routerLink>
      </p>

      <form @submit.prevent="handleSubmit">
        <div class="field">
          <label for="email">Email</label>
          <input
            id="email"
            v-model="form.email"
            type="email"
            placeholder="alex@email.com"
            autocomplete="email"
          />
        </div>

        <div class="field">
          <div class="field-header">
            <label for="password">Пароль</label>
            <a href="#" class="forgot">Забыли пароль?</a>
          </div>
          <input
            id="password"
            v-model="form.password"
            type="password"
            placeholder="••••••••"
            autocomplete="current-password"
          />
        </div>

        <label class="remember">
          <input v-model="form.remember" type="checkbox" />
          <span>Запомнить меня</span>
        </label>

        <button type="submit" class="btn-primary">Войти</button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { RouterLink } from 'vue-router'

const form = reactive({
  email: '',
  password: '',
  remember: false,
})

function handleSubmit() {
  // TODO: connect to your auth API
  console.log('Login:', form)
}

function handleGoogle() {
  // TODO: Google OAuth
  console.log('Google login')
}

function handleFacebook() {
  // TODO: Facebook OAuth
  console.log('Facebook login')
}
</script>

<style scoped>
.login-page {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: var(--bg);
  position: relative;
  overflow: hidden;
  font-family: 'DM Sans', sans-serif;
  color: var(--ink);
}

/* ── Background decor ── */
.bg-decor {
  position: fixed;
  inset: 0;
  pointer-events: none;
}

.circle-1 {
  position: absolute;
  width: 600px;
  height: 600px;
  border-radius: 50%;
  background: radial-gradient(circle, #C8562A0F 0%, transparent 70%);
  top: -200px;
  right: -150px;
}

.circle-2 {
  position: absolute;
  width: 400px;
  height: 400px;
  border-radius: 50%;
  background: radial-gradient(circle, #1A1A1808 0%, transparent 70%);
  bottom: -100px;
  left: -100px;
}

/* ── Card ── */
.card {
  background: var(--white);
  border-radius: 16px;
  padding: 52px 48px;
  width: 100%;
  max-width: 420px;
  position: relative;
  z-index: 1;
  animation: appear 0.5s ease forwards;
  border: 1px solid var(--border);
}

.logo {
  font-family: 'DM Serif Display', serif;
  font-size: 24px;
  color: var(--ink);
  letter-spacing: -0.5px;
  margin-bottom: 36px;
  display: block;
  text-decoration: none;
}

.logo span { color: var(--accent); }

h1 {
  font-family: 'DM Serif Display', serif;
  font-size: 30px;
  letter-spacing: -0.6px;
  margin-bottom: 6px;
}

.sub {
  font-size: 14px;
  color: var(--muted);
  margin-bottom: 32px;
  font-weight: 300;
}

.sub a {
  color: var(--accent);
  text-decoration: none;
  font-weight: 500;
}

.field { margin-bottom: 16px; }

.field-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

label {
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.06em;
  text-transform: uppercase;
  color: var(--muted);
}

.forgot {
  font-size: 12px;
  color: var(--muted);
  text-decoration: none;
  transition: color 0.2s;
}

.forgot:hover { color: var(--ink); }

input[type='email'],
input[type='password'] {
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

input[type='email']::placeholder,
input[type='password']::placeholder { color: #C4C0BA; }

input[type='email']:focus,
input[type='password']:focus {
  border-color: var(--ink);
  background: var(--white);
}

.remember {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 4px 0 20px;
  cursor: pointer;
}

.remember input[type='checkbox'] {
  width: 16px;
  height: 16px;
  accent-color: var(--ink);
  cursor: pointer;
}

.remember span {
  font-size: 13px;
  color: var(--muted);
  font-weight: 300;
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

.social-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}

.btn-social {
  padding: 11px;
  background: transparent;
  color: var(--ink);
  border: 1.5px solid var(--border);
  border-radius: 8px;
  font-family: 'DM Sans', sans-serif;
  font-size: 13px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  transition: border-color 0.2s;
}

.btn-social:hover { border-color: var(--ink); }

@keyframes appear {
  from { opacity: 0; transform: translateY(16px); }
  to   { opacity: 1; transform: translateY(0); }
}

@media (max-width: 480px) {
  .card { padding: 36px 24px; margin: 16px; }
  .social-row { grid-template-columns: 1fr; }
}
</style>
