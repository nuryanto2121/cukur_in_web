package templateemail

const (
	EmailRedem = `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<title>Cukur-in</title>
		</head>
		<body>
			<p>Hai {Name}.</p>
			<p style="text-align: left;">Selamat kamu mendapatkan voucher gratis teguk</p>
			<p style="text-align: left;">Tunjukkan kode voucher berikut di outlet Teguk Indonesia dimanapun</p>
			<p style="text-align: left;">Kode Voucher: {RedemCd}
		
				<br />Tanggal Berlaku: s/d 30 April 2021
	
			</p>
			<p style="text-align: left;">S&amp;K:
		
				<br />1.Voucher berlaku hanya untuk pengguna baru cukur-In
		
				<br />2.Voucher berlaku di seluruh outlet Teguk Indonesia
		
				<br />3.Voucher hanya untuk menu {Menu}
		
				<br />4.Voucher hanya dapat di redeem 1 (satu) kali
		
				<br />5.Voucher tidak dapat di uangkan
	
			</p>
			<p style="text-align: left;">Pilihan outlet Teguk Indonesia terdekat dari lokasimu saat ini</p>
			<ol>
{Outlet}
			</ol>
		</body>
	</html>

	`
)
