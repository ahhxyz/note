/*
 *修改该文件后必须关闭再重新打开页面，只刷新页面是无效的
 */

requirejs.config({
    //By default load any module IDs from js/lib
    baseUrl: 'js/Angular',
    //except, if the module ID starts with "app",
    //load it from the js/app directory. paths
    //config is relative to the baseUrl, and
    //never includes a ".js" extension since
    //the paths config could be for a directory.
    paths: {
        App:'../App',
        Ext: '../Lib'
    }
});

// Start the main app logic.
requirejs(['angular', 'Ext/jquery',  'App/Controller/order'], //这样是加载单个文件
function   (angular,    $,       order) {
    //jQuery, canvas and the app/sub module are all
    //loaded and can be used here now.
});