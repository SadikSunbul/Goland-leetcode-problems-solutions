package _02_AddTwoNumbers

// ListNode yapısı, bağlı liste düğümünü temsil eder.
type ListNode struct {
	Val  int       // Düğümün değeri
	Next *ListNode // Bir sonraki düğüme işaret eder
}

// İki bağlı listeyi toplar ve sonucu yeni bir bağlı liste olarak döndürür.
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Başlangıç için boş bir liste düğümü (dummy node)
	ilkDugum := ListNode{}
	// Sonuç listemizi oluşturacak olan geçici işaretçi
	gecici := &ilkDugum

	// Taşıma değeri (elde)
	tasima := 0

	// Liste elemanlarını toplarken her iki listeyi de döngüyle geziyoruz
	for l1 != nil || l2 != nil {
		// Toplam değeri hesaplamak için geçici bir değişken
		toplam := tasima

		// l1'deki mevcut elemanı ekleyin (varsa)
		if l1 != nil {
			toplam += l1.Val
			l1 = l1.Next // l1'i bir sonraki düğüme geçir
		}

		// l2'deki mevcut elemanı ekleyin (varsa)
		if l2 != nil {
			toplam += l2.Val
			l2 = l2.Next // l2'yi bir sonraki düğüme geçir
		}

		// Yeni taşıma değerini hesapla
		tasima = toplam / 10

		// Yeni düğüm ekle (toplamın son basamağını al)
		gecici.Next = &ListNode{Val: toplam % 10}
		// Geçici işaretçiyi bir sonraki düğüme geçir
		gecici = gecici.Next
	}

	// Eğer taşıma değeri kaldıysa (örneğin 9 + 9 = 18 gibi)
	if tasima != 0 {
		gecici.Next = &ListNode{Val: tasima}
	}

	// Başlangıçtaki dummy düğümünü atıp, gerçek sonucu döndürüyoruz
	return ilkDugum.Next
}
