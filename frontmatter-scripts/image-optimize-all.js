// yarn add i --dev imagemin imagemin-jpegtran imagemin-pngquant
const arguments = process.argv;
const path = require('path');

(async () => {
  if (arguments && arguments.length > 0) {
    const imagemin = (await import('imagemin')).default;
    const imageminJpegtran = (await import('imagemin-jpegtran')).default;
    const imageminPngquant = (await import('imagemin-pngquant')).default;

    const workspaceArg = arguments[2]; // The workspace path
    const folderArg = arguments[3]; // The folder path

    const files = await imagemin([path.join(folderArg, '*.{jpg,png}')], {
      destination: folderArg,
      glob: true,
      plugins: [imageminJpegtran(), imageminPngquant()],
    });

    console.log(`Optimized images: ${files.length}`);
  }
})();
