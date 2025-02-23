package himma

import (
	"io/ioutil"
	"os"
)

func Base() string {
	if _, err := os.Stat("tpl/base.html"); !os.IsNotExist(err) {
		b, _ := ioutil.ReadFile("tpl/base.html")
		return string(b)
	}
	return smartAdminBase()
}

func smartAdminBase() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>
        {{.AppName}}
    </title>
    <meta name="description" content="Page Titile">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no, user-scalable=no, minimal-ui">
    <!-- Call App Mode on ios devices -->
    <meta name="apple-mobile-web-app-capable" content="yes" />
    <!-- Remove Tap Highlight on Windows Phone IE -->
    <meta name="msapplication-tap-highlight" content="no">
    <!-- base css -->
    <link rel="stylesheet" media="screen, print" href="/assets/css/vendors.bundle.css">
    <link rel="stylesheet" media="screen, print" href="/assets/css/app.bundle.css">
    <link rel="stylesheet" media="screen, print" href="/assets/css/notifications/sweetalert2/sweetalert2.bundle.css">
    <link rel="stylesheet" media="screen, print" href="/assets/css/themes/cust-theme-3.css">
    <!-- Place favicon.ico in the root directory -->
    <link rel="apple-touch-icon" sizes="180x180" href="/assets/img/favicon/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/img/favicon.svg">
    <link rel="mask-icon" href="/assets/img/favicon/safari-pinned-tab.svg">
    <link rel="stylesheet" media="screen, print" href="/assets/css/fa-regular.css">
    <link rel="stylesheet" media="screen, print" href="/assets/css/fa-solid.css">
    <link rel="stylesheet" media="screen, print" href="/plugins/bootstrap-table-1.16.0/bootstrap-table.min.css">
    <link rel="stylesheet" media="screen, print" href="/plugins/waitMe-31.10.17/waitMe.min.css">
    <link rel="stylesheet" media="screen, print" href="/plugins/bootstrap-select-1.13.14/css/bootstrap-select.min.css">
    {{ block "css" . }}{{ end }}
</head>
<body class="mod-bg-1 nav-function-fixed">
<!-- BEGIN Page Wrapper -->
<div class="page-wrapper">
    <div class="page-inner">
        <!-- BEGIN Left Aside -->
        <aside class="page-sidebar">
            <div class="page-logo">
                <a href="/" class="page-logo-link press-scale-down d-flex align-items-center position-relative">
                    <img src="/img/logo.png" class="" alt="" aria-roledescription="logo">
                    <span class="page-logo-text mr-1"> {{.AppName}}</span>
                    <span class="position-absolute text-white opacity-50 small pos-top pos-right mr-2 mt-n2"></span>
                </a>
            </div>
            <!-- BEGIN PRIMARY NAVIGATION -->
            <nav id="js-primary-nav" class="primary-nav" role="navigation">
                <div class="info-card">
                    <img src="/img/avatar.png" class="profile-image rounded-circle" alt="{{.AppName}}">
                    <div class="info-card-text">
                        <a href="#" class="d-flex align-items-center text-white">
                            <span class="text-truncate text-truncate-md d-inline-block">
                                {{.Phrase1}}
                            </span>
                        </a>
                        <span class="d-inline-block text-truncate text-truncate-sm">{{.Phrase2}}</span>
                    </div>

                    <img src="/img/cover.png" class="cover" alt="cover">
                </div>

                <ul id="js-nav-menu" class="nav-menu">
                    <li class="nav-title">{{.Description}}</li>
                    {{ block "sidebar" . }}{{ end }}
                </ul>
                <div class="filter-message js-filter-message bg-success-600"></div>
            </nav>
            <!-- END PRIMARY NAVIGATION -->
            <!-- NAV FOOTER -->
            <div class="nav-footer shadow-top">
                <!--
                <a href="#" onclick="return false;" data-action="toggle" data-class="nav-function-minify" class="hidden-md-down">
                    <i class="ni ni-chevron-right"></i>
                    <i class="ni ni-chevron-right"></i>
                </a>
                -->
                <ul class="list-table m-auto nav-footer-buttons">
                    <li>
                        <a href="javascript:void(0);" class="fs-sm" data-toggle="tooltip" data-placement="top" title="" data-original-title="">
                            {{.Year}} &copy; {{.Company}}
                        </a>
                    </li>
                </ul>
            </div>
            <!-- END NAV FOOTER -->
        </aside>
        <!-- END Left Aside -->
        <div class="page-content-wrapper">
            <!-- BEGIN Page Header -->
            <header class="page-header" role="banner">
                <!-- we need this logo when user switches to nav-function-top -->
                <div class="page-logo">
                    <a href="#" class="page-logo-link press-scale-down d-flex align-items-center position-relative">
                        <img src="/img/logo.png" alt="{{.AppName}}" aria-roledescription="logo">
                        <span class="page-logo-text mr-1">{{.AppName}}</span>
                        <span class="position-absolute text-white opacity-50 small pos-top pos-right mr-2 mt-n2"></span>
                        <i class="fal fa-angle-down d-inline-block ml-1 fs-lg color-primary-300"></i>
                    </a>
                </div>
                <div class="hidden-md-down dropdown-icon-menu position-relative">
                    <a href="#" class="header-btn btn js-waves-off" data-action="toggle" data-class="nav-function-minify" title="Hide Navigation">
                        <i class="ni ni-menu"></i>
                    </a>
                </div>
                <!-- DOC: mobile button appears during mobile width -->
                <div class="hidden-lg-up">
                    <a href="#" class="header-btn btn press-scale-down" data-action="toggle" data-class="mobile-nav-on">
                        <i class="ni ni-menu"></i>
                    </a>
                </div>
            </header>
            <!-- END Page Header -->
            <!-- BEGIN Page Content -->
            <!-- the #js-page-content id is needed for some plugins to initialize -->
            <main id="js-page-content" role="main" class="page-content">
                <!-- Content ### -->
                {{ block "content" . }}{{ end }}
            </main>
            <!-- this overlay is activated only when mobile menu is triggered -->
            <div class="page-content-overlay" data-action="toggle" data-class="mobile-nav-on"></div> <!-- END Page Content -->

            <footer class="page-footer" role="contentinfo">
                <div class="d-flex align-items-center flex-1 text-muted">
                    <span class="hidden-md-down ">
                        <a href="{{.Url}}" class="text-primary fw-500" title="gotbootstrap.com" target="_blank"></a>
                    </span>
                </div>
                <div>
                    <ul class="list-table m-0">
                        <li><a href="{{.Url}}" class="text-secondary">{{.AppName}} v{{.Version}}</a></li>
                    </ul>
                </div>
            </footer>

            <!-- END Page Footer -->
            <!-- BEGIN Color profile -->
            <!-- this area is hidden and will not be seen on screens or screen readers -->
            <!-- we use this only for CSS color refernce for JS stuff -->
            <p id="js-color-profile" class="d-none">
                <span class="color-primary-50"></span>
                <span class="color-primary-100"></span>
                <span class="color-primary-200"></span>
                <span class="color-primary-300"></span>
                <span class="color-primary-400"></span>
                <span class="color-primary-500"></span>
                <span class="color-primary-600"></span>
                <span class="color-primary-700"></span>
                <span class="color-primary-800"></span>
                <span class="color-primary-900"></span>
                <span class="color-info-50"></span>
                <span class="color-info-100"></span>
                <span class="color-info-200"></span>
                <span class="color-info-300"></span>
                <span class="color-info-400"></span>
                <span class="color-info-500"></span>
                <span class="color-info-600"></span>
                <span class="color-info-700"></span>
                <span class="color-info-800"></span>
                <span class="color-info-900"></span>
                <span class="color-danger-50"></span>
                <span class="color-danger-100"></span>
                <span class="color-danger-200"></span>
                <span class="color-danger-300"></span>
                <span class="color-danger-400"></span>
                <span class="color-danger-500"></span>
                <span class="color-danger-600"></span>
                <span class="color-danger-700"></span>
                <span class="color-danger-800"></span>
                <span class="color-danger-900"></span>
                <span class="color-warning-50"></span>
                <span class="color-warning-100"></span>
                <span class="color-warning-200"></span>
                <span class="color-warning-300"></span>
                <span class="color-warning-400"></span>
                <span class="color-warning-500"></span>
                <span class="color-warning-600"></span>
                <span class="color-warning-700"></span>
                <span class="color-warning-800"></span>
                <span class="color-warning-900"></span>
                <span class="color-success-50"></span>
                <span class="color-success-100"></span>
                <span class="color-success-200"></span>
                <span class="color-success-300"></span>
                <span class="color-success-400"></span>
                <span class="color-success-500"></span>
                <span class="color-success-600"></span>
                <span class="color-success-700"></span>
                <span class="color-success-800"></span>
                <span class="color-success-900"></span>
                <span class="color-fusion-50"></span>
                <span class="color-fusion-100"></span>
                <span class="color-fusion-200"></span>
                <span class="color-fusion-300"></span>
                <span class="color-fusion-400"></span>
                <span class="color-fusion-500"></span>
                <span class="color-fusion-600"></span>
                <span class="color-fusion-700"></span>
                <span class="color-fusion-800"></span>
                <span class="color-fusion-900"></span>
            </p>
            <!-- END Color profile -->
        </div>
    </div>
</div>
<!-- END Page Wrapper -->
<!-- BEGIN Quick Menu -->
<!-- to add more items, please make sure to change the variable '$menu-items: number;' in your _page-components-shortcut.scss -->
<nav class="shortcut-menu d-none d-sm-block">
    <input type="checkbox" class="menu-open" name="menu-open" id="menu_open" />
    <label for="menu_open" class="menu-open-button ">
        <span class="app-shortcut-icon d-block"></span>
    </label>
    <a href="#" class="menu-item btn" data-toggle="tooltip" data-placement="left" title="Scroll Top">
        <i class="fal fa-arrow-up"></i>
    </a>
    <a href="#" class="menu-item btn" data-action="app-fullscreen" data-toggle="tooltip" data-placement="left" title="Full Screen">
        <i class="fal fa-expand"></i>
    </a>
    <a href="#" class="menu-item btn" data-action="app-print" data-toggle="tooltip" data-placement="left" title="Print page">
        <i class="fal fa-print"></i>
    </a>
</nav>
<!-- END Quick Menu -->
<script src="/assets/js/vendors.bundle.js"></script>
<script src="/assets/js/app.bundle.js"></script>
<script src="/assets/js/notifications/sweetalert2/sweetalert2.bundle.js"></script>
<script src="/plugins/waitMe-31.10.17/waitMe.min.js"></script>
<script src="/plugins/progressbar.js-1.0.1/progressbar.min.js"></script>
<script src="/plugins/moment-2.27/moment-with-locales.min.js"></script>
<script src="/plugins/moment-2.27/moment-timezone-with-data.min.js"></script>
<script src="/plugins/js.cookie-2.2.1/js.cookie-2.2.1.min.js"></script>
<script src="/plugins/jquery-validation-1.19.2/jquery.validate.min.js"></script>
<script src="/plugins/jquery-mask-1.14.16/jquery.mask.min.js"></script>
<script src="/plugins/holder-2.9/holder.min.js"></script>
<script src="/plugins/bootstrap-table-1.16.0/bootstrap-table.min.js"></script>
<script src="/plugins/bootstrap-table-1.16.0/extensions/bootstrap-table-export.min.js"></script>
<script src="/plugins/bootstrap-table-1.16.0/tableExport.min.js"></script>
<script src="/plugins/bootstrap-table-1.16.0/extensions/bootstrap-table-cookie.min.js"></script>
<script src="/plugins/bootstrap-select-1.13.14/js/bootstrap-select.min.js"></script>
{{ block "script" . }}{{ end }}
</body>
</html>`
}
