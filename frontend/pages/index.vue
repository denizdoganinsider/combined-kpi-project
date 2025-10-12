<template>
  <section class="inner_content new_index background_10" fetchpriority="high">
        <div id="media_v4" class="media discover" bis_skin_checked="1">
          <div class="column_wrapper" bis_skin_checked="1">
            <div class="content_wrapper wrap" bis_skin_checked="1">
              <div class="title" bis_skin_checked="1">
                <h2>Welcome.</h2>
                <h3>Millions of movies, TV shows and people to discover. Explore now.</h3>
              </div>

              <div class="search" bis_skin_checked="1">
                <form id="inner_search_form" action="/search" method="get" accept-charset="utf-8">
                  <label class="inner_search_wrapper">
                    <span class="k-input k-textbox k-input-solid k-input-md k-rounded-md" style=""><input dir="auto" id="inner_search_v4" name="query" type="text" autocorrect="off" autofill="off" autocomplete="off" spellcheck="false" placeholder="Search for a movie, tv show, person......" value="" data-role="textbox" aria-disabled="false" class="k-input-inner" inputmode="text" style="width: 100%;"></span>
                  </label>

                  <input type="submit" value="Search" data-role="button" class="k-button k-button-md k-rounded-md k-button-solid k-button-solid-base" role="button" aria-disabled="false" tabindex="0">
                </form>
              </div>
            </div>
          </div>
        </div>
      </section>
  <div v-if="isAuthenticated" class="p-8 api-test">
    <h1 class="text-2xl font-bold text-gray-800 mb-4">API Response:</h1>
    <p class="text-lg text-green-600">{{ message }}</p>
  </div>
  <div v-else>
    <p>You're redirecting...</p>
  </div>
</template>

<script setup>
  definePageMeta({
    requiresAuth: true
  })

import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

/* i commented this line for quick testing */
/* const isAuthenticated = process.client ? localStorage.getItem('auth_token') !== null : false; */
const isAuthenticated = true;

const message = ref('Loading...')
const router = useRouter()

onMounted(async () => {
  if (!isAuthenticated) {
    router.push('/login')
  } else {
    /*
    try {
      const res = await axios.get('http://localhost:8080/api/v1/hello')
      message.value = res.data.message
    } catch (error) {
      message.value = 'Error occurred: ' + error.message
    }
      */
  }
})
</script>

<style scoped>
.api-test {
  display: none;
}
section.new_index.background_10 {
    border-width: 0;
    border-style: solid;
    border-color: #efefef;
    background-position: 50% 50%;
    flex-wrap: wrap;
    align-items: flex-start;
    align-content: flex-start;
    box-sizing: border-box;
    width: 100%;
    display: flex;
    justify-content: center;
    min-height: 300px;
    height: calc(100vh/2.5);
    max-height: 360px;
    background-color: rgba(0,0,0,.2);
    background-size: cover;
    background-repeat: no-repeat;
    color: #fff;
    min-width: 100%;
    max-width: var(--maxPrimaryPageWidth);
    background-image: image-set(url(https://media.themoviedb.org/t/p/w1920_and_h600_multi_faces_filter\(duotone,00192f,00baff\)/npD65vPa4vvn1ZHpp3o05A5vdKT.jpg)1x,url(https://media.themoviedb.org/t/p/w3840_and_h1200_multi_faces_filter\(duotone,00192f,00baff\)/npD65vPa4vvn1ZHpp3o05A5vdKT.jpg)2x);
}

section.new_index div.media {
    height: 100%;
    display: flex;
    align-content: center;
    align-items: center;
    justify-content: center;
}

div.column_wrapper, div.column_wrapper div.column {
    width: 100%;
}

div.column_wrapper, div.column_wrapper div.content_wrapper {
    display: flex;
    align-items: flex-start;
    align-content: flex-start;
}

div.column_wrapper {
    justify-content: center;
    min-width: 100%;
}

div.media.discover div.content_wrapper {
    flex-wrap: wrap;
}

div.column_wrapper div.content_wrapper {
    max-width: 1400px;
    width: 100%;
    padding: 30px 40px;
}

div.media.discover div.content_wrapper>div.title {
    width: 100%;
    margin-bottom: 20px;
}

section.new_index form {
    margin-top: 30px;
    position: relative;
    top: 0;
    left: 0;
}

section.new_index form input[type=submit] {
    height: 46px;
    padding: 10px 26px;
    border: 0;
    background: linear-gradient(to right, rgba(30, 213, 169, 1) 0%, rgba(1, 180, 228, 1) 100%);
    border-radius: 30px;
    position: absolute;
    top: 1px;
    right: -1px;
    z-index: 2;
    color: #fff;
}

section.new_index form input[type=text] {
    width: 100%;
    height: 46px;
    line-height: 46px;
    font-size: 1.1em;
    color: rgba(0, 0, 0, .5);
    border: 0;
    border-radius: 30px;
    padding: 10px 20px;
}

section.new_index form label.inner_search_wrapper>span {
    border-radius: 30px;
}

div.media.discover div.content_wrapper>div.title h2 {
    font-size: 3em;
    font-weight: 700;
    line-height: 1;
}

section.new_index h3 {
    font-size: 2em;
    font-weight: 600;
    margin: 0;
}

section.new_index div.search, section.new_index div.title {
    width: 100%;
}
</style>
