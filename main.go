// index.js

require('dotenv').config();
const axios = require('axios');
const yargs = require('yargs/yargs');
const { hideBin } = require('yargs/helpers');

// コマンドライン引数の設定
const argv = yargs(hideBin(process.argv))
  .option('city', {
    alias: 'c',
    description: '天気を取得する都市名',
    type: 'string',
    default: ''
  })
  .option('location', {
    alias: 'l',
    description: '天気を取得する都市名 (完全名)',
    type: 'string',
    default: 'Tokyo'
  })
  .help()
  .alias('help', 'h')
  .argv;

// メイン関数
async function main() {
  // 場所選択
  const locations = {
    'tokyo': 'Tokyo',
    'la': 'Los Angeles',
    'osaka': 'Osaka',
    'hokkaido': 'Sapporo'
  };

  let locationName = argv.location;

  // city引数が指定されていれば、それを使用
  if (argv.city) {
    if (locations[argv.city]) {
      locationName = locations[argv.city];
    } else {
      console.log(`未知の場所です: ${argv.city}`);
      console.log('利用可能な場所: ' + Object.entries(locations).map(([k, v]) => `${k} (${v})`).join(' '));
      return;
    }
  }

  try {
    // 指定された場所の天気を取得
    const weather = await getWeather(locationName);

    // 日本のタイムゾーンを設定
    const now = new Date();
    const jstTime = new Date(now.getTime() + (9 * 60 * 60 * 1000));
    
    // 結果を表示
    console.log(`日時: ${jstTime.toISOString().replace('T', ' ').substring(0, 19)}`);
    console.log(`場所: ${weather.name}`);
    if (weather.weather && weather.weather.length > 0) {
      console.log(`天気: ${weather.weather[0].description}`);
    }
    console.log(`気温: ${weather.main.temp.toFixed(1)}°C`);
  } catch (error) {
    console.error('天気情報の取得に失敗しました:', error.message);
  }
}

// 指定された場所の天気情報を取得する関数
async function getWeather(location) {
  // 環境変数からAPIキーを取得
  const apiKey = process.env.OPENWEATHERMAP_API_KEY;
  if (!apiKey) {
    throw new Error('APIキーが設定されていません。.envファイルまたは環境変数にOPENWEATHERMAP_API_KEYを設定してください');
  }
  
  // OpenWeatherMap APIのエンドポイント
  const url = `https://api.openweathermap.org/data/2.5/weather?q=${location}&units=metric&appid=${apiKey}`;

  // HTTPリクエスト
  const response = await axios.get(url);
  return response.data;
}

// プログラム実行
main().catch(err => {
  console.error(err);
  process.exit(1);
});