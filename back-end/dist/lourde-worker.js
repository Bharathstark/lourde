importScripts("/precache-manifest.649a426fd8df4dcec37585b0fa2a0979.js", "https://storage.googleapis.com/workbox-cdn/releases/4.3.1/workbox-sw.js");

workbox.setConfig({
  debug: true,
});

workbox.core.setCacheNameDetails({ prefix: "lourde" });

self.addEventListener("message", (event) => {
  if (event.data && event.data.type === "SKIP_WAITING") {
    self.skipWaiting();
  }
});

self.addEventListener('install', event => {
  self.skipWaiting();
});

self.addEventListener('activate', () => {
  self.clients.claim();
  
});

self.__precacheManifest = [].concat(self.__precacheManifest || []);
workbox.precaching.precacheAndRoute(self.__precacheManifest, {});
workbox.routing.registerNavigationRoute("/index.html");
/**
 * The workboxSW.precacheAndRoute() method efficiently caches and responds to
 * requests for URLs in the manifest.
 * See https://goo.gl/S9QRab
 */



workbox.routing.registerRoute(
  /\.(?:png|gif|jpg|JPG|jpeg|svg|webp)$/,
  new workbox.strategies.CacheFirst({
    cacheName: "images",
    plugins: [
      new workbox.expiration.Plugin({
        maxAgeSeconds: 30 * 24 * 60 * 60, // 30 Days
      }),
    ],
  })
);

workbox.routing.registerRoute(
  /\.(?:js|css|html)$/,
  new workbox.strategies.StaleWhileRevalidate({
    cacheName: "files",
    plugins: [
      new workbox.expiration.Plugin({
        maxAgeSeconds: 30 * 24 * 60 * 60, // 30 Days
      }),
    ],
  })
);

workbox.routing.registerRoute(
  ({ url }) => url.pathname.startsWith("/api/v1"),
  new workbox.strategies.NetworkFirst({
    cacheName: "api",
    plugins: [
      new workbox.expiration.Plugin({
        maxAgeSeconds: 30 * 24 * 60 * 60, // 30 Days
      }),
      new workbox.cacheableResponse.Plugin({
        statuses: [0, 200],
    })
    ],
  })
);

workbox.routing.registerRoute(
  ({ url }) => url.origin.startsWith("https://fonts.googleapis.com"),
  new workbox.strategies.StaleWhileRevalidate({
    cacheName: "google-fonts-stylesheets",
    plugins: [
      new workbox.expiration.Plugin({
        maxAgeSeconds: 30 * 24 * 60 * 60, // 30 Days
      }),
    ],
  })
);


workbox.routing.registerRoute(
  ({ url }) => url.origin.startsWith("https://fonts.gstatic.com"),
  new workbox.strategies.CacheFirst({
    cacheName: "google-fonts-webfonts",
    plugins: [
      new workbox.expiration.Plugin({
        maxAgeSeconds: 30 * 24 * 60 * 60, // 30 Days
      }),
    ],
  })
);

workbox.routing.registerRoute(
  ({ url }) => url.origin.startsWith("https://www.google.com/maps/embed"),
  new workbox.strategies.CacheFirst({
    cacheName: "google-maps",
    plugins: [
      new workbox.expiration.Plugin({
        maxAgeSeconds: 30 * 24 * 60 * 60, // 30 Days
      }),
    ],
  })
);



