<template>
  <div>
    <div>
      <h2>
        <a href="/" class="header-link">
          <i class="pi pi-spin pi-cog" style="margin-right: 0.5rem"></i> Web
          Page Analyzer
        </a>
      </h2>
    </div>
    <Divider />

    <div class="form">
      <div class="card flex justify-center">
        <div class="input-group">
          <InputText
            v-model="url"
            type="text"
            placeholder="Enter URL here"
            style="width: 50rem"
            :invalid="!url.trim()"
          />
          <Button
            label="Analyze"
            severity="info"
            @click="analyze"
            :disabled="!url.trim() || isLoading"
          />
          <Button
            label="Clear"
            severity="danger"
            @click="clearOutput"
            :disabled="isLoading"
          />
        </div>
      </div>

      <div v-if="isLoading" class="loading-container">
        <ProgressSpinner />
      </div>

      <div>
        <div class="section" v-if="isLoaded">
          <div class="header">Results</div>
          <div class="section horizontal">
            <div class="field">
              <label>Title</label>
              <InputText
                id="address_name"
                v-model="result.title"
                disabled="True"
              />
            </div>

            <div class="field">
              <label>HTML Version</label>
              <InputText
                v-model="result.html_version"
                :value="
                  result.html_version ? result.html_version : 'Not Specified'
                "
                disabled="True"
              />
            </div>

            <div class="field">
              <label>Internal Links</label>
              <InputText v-model="result.internal_links" disabled="True" />
            </div>

            <div class="field">
              <label>External Links</label>
              <InputText v-model="result.external_links" disabled="True" />
            </div>

            <div class="field">
              <label>Login Form Present</label>
              <InputText
                v-model="result.login_form"
                :value="result.login_form ? 'Yes' : 'No'"
                disabled="True"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import Button from "primevue/button";
import Divider from "primevue/divider";
import InputText from "primevue/inputtext";
import Message from "primevue/message";
import ProgressSpinner from "primevue/progressspinner";
import { useToast } from "primevue/usetoast";

const toast = useToast();
const url = ref("");
const result = ref({
  title: "",
  html_version: "Not specified",
  internal_links: 0,
  external_links: 0,
  inaccessible_links: 0,
  login_form: false,
});
const isLoading = ref<boolean>(false);
const isLoaded = ref<boolean>(false);

const analyze = async () => {
  isLoaded.value = false;
  isLoading.value = true;
  result.value = {
    title: "",
    html_version: "Not specified",
    internal_links: 0,
    external_links: 0,
    inaccessible_links: 0,
    login_form: false,
  };

  try {
    const response = await fetch("http://localhost:8080/analyze", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url: url.value }),
    });
    const resp = await response.json();

    if (response.ok) {
      result.value = resp.data;
      isLoaded.value = true;
    } else {
      toast.add({ severity: "error", detail: resp.error, life: 3000 });
    }
  } catch (err) {
    toast.add({ severity: "error", detail:  "Analyzing Failed.", life: 3000 });
  } finally {
    isLoading.value = false;
  }
};

const clearOutput = () => {
  url.value = "";
  result.value = {
    title: "",
    html_version: "Not specified",
    internal_links: 0,
    external_links: 0,
    inaccessible_links: 0,
    login_form: false,
  };
  isLoading.value = false;
  isLoaded.value = false;
};
</script>

<style>
a.header-link {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: inherit;
}
.p-divider.p-divider-horizontal:before {
  border-top: 1px solid rgb(107, 106, 106) !important;
}
.background-card {
  background-color: #f4f4f4; /* Light background for the card */
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1); /* Subtle shadow */
  border-radius: 8px; /* Rounded corners */
}

.card-content p {
  margin-bottom: 12px;
  font-size: 1rem;
}

.card-content p strong {
  font-weight: 600;
}
.loading-container {
  display: flex;
  justify-content: center;
  align-items: center;

  flex-direction: column;
}
.input-group {
  display: flex;
  justify-content: center;
  margin-bottom: 1rem;
  gap: 1rem;
}

.form {
  display: flex;
  border: 0.1rem solid #ddd;
  padding: 1rem;
  border-radius: 1rem;
  margin: 1rem auto;
  background-color: #f9f9f9;
  box-shadow: 0 1rem 1rem rgba(0, 0, 0, 0.1);
  flex-direction: column;
}

.form .header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.form .section .header {
  font-size: 1.5em;
  font-weight: bold;
  margin-bottom: 1rem;
  color: #333;
}

.form .edit-button {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 1rem;
}

.form .section {
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
  gap: 1rem;
}

.form .section.horizontal {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
}

.form .field {
  display: flex;
  flex-direction: column;
  label {
    font-weight: bold;
    margin-bottom: 1rem;
    color: #555;
  }
}

.form .submit {
  display: flex;
  justify-content: end;
  align-items: end;
  margin-top: 1rem;
  gap: 1rem;
}
</style>
