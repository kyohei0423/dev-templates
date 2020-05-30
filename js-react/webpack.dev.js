const webpack = require('webpack')
const merge = require('webpack-merge')
const common = require('./webpack.common.js')

module.exports = merge(common, {
  mode: 'development',
  devtool: 'inline-source-map', // デバッグ
  // web server
  devServer: {
    contentBase: './dist',
    host: '0.0.0.0',
    port: '83',
    hot: true,
    historyApiFallback: true, // 404のときにindex.htmlを返す
    allowedHosts: ['lvh.me'],
  },
  plugins: [new webpack.HotModuleReplacementPlugin()],
})
