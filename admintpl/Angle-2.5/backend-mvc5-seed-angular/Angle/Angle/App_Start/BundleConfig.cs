using System.Web;
using System.Web.Optimization;

namespace Angle
{
    public class BundleConfig
    {
        // For more information on bundling, visit http://go.microsoft.com/fwlink/?LinkId=301862
        public static void RegisterBundles(BundleCollection bundles)
        {
            // Application Scripts
            bundles.Add(new ScriptBundle("~/bundles/appScripts")
                .IncludeDirectory("~/Scripts/app", "*.js", true));
            
            // Base Scripts (not lazyloaded)
            bundles.Add(new ScriptBundle("~/bundles/baseScripts").Include(
              "~/Vendor/modernizr/modernizr.js",
              "~/Vendor/jquery/dist/jquery.js",
              "~/Vendor/angular/angular.js",
              "~/Vendor/angular-route/angular-route.js",
              "~/Vendor/angular-cookies/angular-cookies.js",
              "~/Vendor/angular-animate/angular-animate.js",
              "~/Vendor/angular-touch/angular-touch.js",
              "~/Vendor/angular-ui-router/release/angular-ui-router.js",
              "~/Vendor/ngstorage/ngStorage.js",
              "~/Vendor/angular-ui-utils/ui-utils.js",
              "~/Vendor/angular-sanitize/angular-sanitize.js",
              "~/Vendor/angular-resource/angular-resource.js",
              "~/Vendor/angular-translate/angular-translate.js",
              "~/Vendor/angular-translate-loader-url/angular-translate-loader-url.js",
              "~/Vendor/angular-translate-loader-static-files/angular-translate-loader-static-files.js",
              "~/Vendor/angular-translate-storage-local/angular-translate-storage-local.js",
              "~/Vendor/angular-translate-storage-cookie/angular-translate-storage-cookie.js",
              "~/Vendor/oclazyload/dist/ocLazyLoad.js",
              "~/Vendor/angular-bootstrap/ui-bootstrap-tpls.js",
              "~/Vendor/angular-loading-bar/build/loading-bar.js",
              "~/Vendor/angular-dynamic-locale/dist/tmhDynamicLocale.js"
            ));

            bundles.Add(new StyleBundle("~/bundles/appStyles")
                    .Include("~/Content/app/app.css"));
            
            bundles.Add(new StyleBundle("~/bundles/bootstrapStyles")
                    .Include("~/Content/app/bootstrap.css"));
        }
    }
}
