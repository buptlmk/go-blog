const {override, fixBabelImports, addLessLoader} = require('customize-cra');
module.exports = override(
    fixBabelImports('import', {
        libraryName: 'antd',
        libraryDirectory: 'es',
        style: true,
    }),
    addLessLoader({
        javascriptEnabled: true,
        // strictMath: true,
        noIeCompat: true,
        // modules: true,
        // modifyVars: { '@primary-color': '#1DA57A' },
        // importLoaders: true,
        // localIdentName: '[local]--[hash:base64:5]' // 自定义 CSS Modules 的 localIdentName
    }),
);
