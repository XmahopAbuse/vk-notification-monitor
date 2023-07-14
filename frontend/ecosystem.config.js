module.exports = {
  script: 'serve',
  autorestart: true,
  name: 'vk-monitor',
  env: {
    PM2_SERVE_PATH: 'build',
    PM2_SERVE_PORT: 5000,
  },
};
