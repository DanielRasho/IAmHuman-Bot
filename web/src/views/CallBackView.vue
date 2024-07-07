<template>
  <div class="body">
    <div class="card">
      <img src="../assets/dragon.png" alt="" />
      <h1>{{ header }}</h1>

      <p v-if="requestSuccess">
        Se te asigno a
        <span>{{ class_subject }}</span>
        seccion <span> {{ class_section }} </span> del server.
      </p>
    </div>
  </div>
</template>
<script setup>
import { onMounted, ref } from 'vue'
import { APISettings } from '@/utils/apiConfig'
import { useRoute } from 'vue-router'

const route = useRoute()

const requestSuccess = ref(false)

const header = ref('Procesando...')
const class_subject = ref('...')
const class_section = ref('...')

onMounted(async () => {
  const url = APISettings.baseURL + 'assign-role/success'
  const searchParams = new URLSearchParams(window.location.search)
  const code = searchParams.get('code')
  const invitationId = searchParams.get('state')
  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: APISettings.Headers,
      body: JSON.stringify({
        code: code,
        invitation_id: invitationId
      })
    })
    const data = await response.json()
    if (!response.ok) throw Error(data.message)
    await getInvitationData()
    header.value = 'Exito!'
    requestSuccess.value = true
  } catch (err) {
    header.value = 'Ups... Error'
  }
})

async function getInvitationData() {
  const searchParams = new URLSearchParams(window.location.search)
  const id = searchParams.get('state')
  const url = APISettings.baseURL + 'assign-role/?invitation_id=' + id
  try {
    const response = await fetch(url, {
      method: 'GET'
    })
    const data = await response.json()
    if (!response.ok) throw Error(data.message)
    class_subject.value = data.invitation_data.class
    class_section.value = data.invitation_data.section
  } catch (err) {
    alert('Cannot fetch invitation data\n' + err)
  }
}
</script>
<style>
.body {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.card {
  background-color: var(--grey-01);

  display: flex;
  flex-direction: column;
  align-items: center;
  width: 40%;
  max-width: 600px;

  padding: 1rem 2rem;
  border-radius: 0.8rem;
  box-shadow:
    rgba(0, 0, 0, 0.25) 0px 13px 27px -5px,
    rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
}

.card > * {
  margin: 1rem 0;
}

.card > h1 {
  font-weight: bolder;
  font-size: 2rem;
  letter-spacing: 0.1ch;
  background: linear-gradient(to right, var(--green-light), var(--blue));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.card > img {
  border-radius: 50%;
  width: 15ch;
}

.card > p {
  text-align: center;
}

.card > p > span {
  font-weight: bold;
  color: var(--white);
}

@media only screen and (max-width: 700px) {
  .card {
    background-color: transparent;
    box-shadow: none;
    width: 100%;
  }
}
</style>
