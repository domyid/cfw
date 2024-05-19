import wasm from '../pkg/main.wasm';

addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request));
});

async function handleRequest(request) {
  await wasm();
  const url = GOOGLE_CLOUD_FUNCTION_URL;
  const response = await window.handleRequest(request,url);
  return new Response(await response.text(), {
    headers: { "Content-Type": "application/json" }
  });
}
