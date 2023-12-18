import path from 'path';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import CopyWebpackPlugin from 'copy-webpack-plugin';
import HtmlWebpackTagsPlugin from 'html-webpack-tags-plugin';


export default {
  entry: path.resolve(__dirname, './src/main.js'),
  output: {
    filename: 'js/[name]-[hash:6].js',
    path: path.resolve(__dirname, './dist')
  },
  module: {
    rules: [
      {
        test: /\.(?:js|mjs|cjs)$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: [
              ['@babel/preset-env', { targets: "defaults" }]
            ]
          }
        }
      },
      {
        test: /\.html$/,
        use: [
          {
            loader: 'html-loader'
          }
        ]
      },
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, './public/index.html'),
      filename: path.resolve(__dirname, './dist/index.html'),
      inject: 'body'
    }),
    new CopyWebpackPlugin({
      patterns: [
        { context: path.resolve(__dirname, 'public'), from: '**.js', to: 'js', info: { minimized: true } },
        { context: path.resolve(__dirname, 'public'), from: '**.wasm', to: '', info: { minimized: true } },
      ],
    }),
    new HtmlWebpackTagsPlugin({
      scripts: [
        'js/wasm_exec.js',
      ],
      append: false,
      hash: true,
    }),
  ],
  devServer: {
    port: 8000,
    allowedHosts: 'all',
  }
};
