#!/bin/bash
# set -euo pipefail

# Fungsi untuk menampilkan teks berwarna
log_info() {
    echo -e "\033[0;34m[INFO]\033[0m $1"
}

log_success() {
    echo -e "\033[0;32m[SUCCESS]\033[0m $1"
}

log_error() {
    echo -e "\033[0;31m[ERROR]\033[0m $1"
}

# 1. Validasi: Pastikan ini adalah folder Git
if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    log_error "Folder ini bukan repositori Git!"
    exit 1
fi

show_help() {
    echo "======================================================="
    echo "       🚀 PROJECT GIT HELPER UTILITY (CLI MODE)        "
    echo "======================================================="
    echo "Penggunaan: ./git-helper.sh [FLAG]"
    echo ""
    echo "Pilihan Flag:"
    echo "  -c, --checkout  Pindah branch berbasis nomor + otomatis pull"
    echo "  -s, --sync      Sinkronisasi branch saat ini dengan main/master"
    echo "  -m, --commit    Instan commit (add ., commit, & push) aman"
    echo "  -h, --help      Tampilkan menu bantuan ini"
    echo "======================================================="
}

smart_checkout() {
    log_info "Mengambil daftar branch lokal..."

    # Mengambil daftar branch secara dinamis ke dalam array
    IFS=$'\n' read -r -d '' -a branches < <(
        git branch --format='%(refname:short)'
        printf '\033'
    )

    if [ ${#branches[@]} -eq 0 ]; then
        log_error "Tidak ada branch lokal yang ditemukan!"
        return
    fi

    echo "======================================================="
    echo "       📝 SILAKAN PILIH BRANCH TUJUAN KAMU             "
    echo "   (Ketik nomor branch lalu ENTER, atau angka acak untuk batal) "
    echo "======================================================="
    echo ""

    PS3="Masukkan pilihan nomor [1-${#branches[@]}]: "

    select target_branch in "${branches[@]}"; do
        if [ -n "$target_branch" ]; then
            echo ""
            read -p "⚠️ Kamu memilih branch '$target_branch'. Lanjutkan checkout? (y/n): " confirm_checkout

            if [[ ! "$confirm_checkout" =~ ^[Yy]$ ]]; then
                log_info "Checkout ke '$target_branch' dibatalkan."
                break
            fi

            log_info "Mencoba pindah ke branch '$target_branch'..."

            if git checkout "$target_branch" &&
                git pull origin "$target_branch"; then
                log_success "Branch berhasil diperbarui!"
                break
            else
                log_error "Gagal memindahkan atau memperbarui branch. Periksa konflik/koneksi Anda."
                exit 1
            fi
        else
            log_info "Pemilihan branch dibatalkan."
            break
        fi
    done

    unset PS3
}

sync_with_main() {
    local current_branch
    current_branch=$(git branch --show-current)

    local main_branch="main"

    if ! git show-ref --verify --quiet refs/heads/main; then
        if git show-ref --verify --quiet refs/heads/master; then
            main_branch="master"
        else
            log_error "Tidak mendeteksi branch 'main' atau 'master' di lokal."
            return
        fi
    fi

    log_info "Branch utama terdeteksi: $main_branch"
    log_info "Alur: Pindah ke $main_branch -> Pull -> Kembali ke $current_branch -> Merge $main_branch"

    read -p "Lanjutkan sinkronisasi? (y/n): " confirm

    if [[ ! "$confirm" =~ ^[Yy]$ ]]; then
        log_info "Sinkronisasi dibatalkan."
        return
    fi

    if git checkout "$main_branch" &&
        git pull origin "$main_branch" &&
        git checkout "$current_branch" &&
        git merge "$main_branch"; then
        log_success "Branch '$current_branch' berhasil disinkronkan dengan '$main_branch'!"
    else
        log_error "Terjadi konflik atau kegagalan saat melakukan merge! Selesaikan secara manual."
        exit 1
    fi
}

quick_commit() {
    local current_branch
    current_branch=$(git branch --show-current)

    log_info "Menjalankan Instan Commit di branch: '$current_branch'"

    if git diff-index --quiet HEAD --; then
        log_info "Mencoba cek berkas baru (untracked)..."

        if [ -z "$(git status --porcelain)" ]; then
            log_success "Repositori bersih. Tidak ada yang perlu di-commit!"
            return
        fi
    fi

    read -p "Masukkan pesan commit (commit message): " commit_msg

    if [ -z "$commit_msg" ]; then
        log_error "Pesan commit tidak boleh kosong! Proses dibatalkan."
        return
    fi

    log_info "Mengeksekusi git add . dan git commit..."

    if git add . &&
        git commit -m "$commit_msg" &&
        log_info "Mengirim perubahan ke origin/$current_branch..." &&
        git push origin "$current_branch"; then
        log_success "Berhasil melakukan pencatatan dan unggah data!"
    else
        log_error "Proses Git Instan Commit gagal! Periksa aturan commitlint atau konflik remote di atas."
        exit 1
    fi
}

# ---- PARSING ARGUMEN / FLAGS ----

# Jika tidak ada argumen sama sekali (hanya ngetik ./git-helper.sh)
if [ $# -eq 0 ]; then
    show_help
    exit 0
fi

# Membaca flag yang dimasukkan pengguna
case "$1" in
    -c|--checkout)
        smart_checkout
        ;;
    -s|--sync)
        sync_with_main
        ;;
    -m|--commit)
        quick_commit
        ;;
    -h|--help)
        show_help
        ;;
    *)
        log_error "Flag '$1' tidak dikenali!"
        show_help
        exit 1
        ;;
esac