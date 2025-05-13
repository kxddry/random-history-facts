package factmatcher

import (
    "regexp"
    "strings"
    "unicode"
)

type Fact_Matcher struct{}

// normalize делает строку нижнего регистра, убирает пунктуацию и лишние пробелы
func (f Fact_Matcher) Normalize(s string) string {
    // Приводим к нижнему регистру
    s = strings.ToLower(s)
    
    // Удаляем пунктуацию
    var sb strings.Builder
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
            sb.WriteRune(r)
        }
    }
    
    // Удаляем лишние пробелы
    space := regexp.MustCompile(`\s+`)
    clean := space.ReplaceAllString(sb.String(), " ")
    
    return strings.TrimSpace(clean)
}

// tokenize разбивает строку на уникальные слова
func tokenize(s string) map[string]struct{} {
    words := strings.Fields(s)
    tokens := make(map[string]struct{}, len(words))
    for _, word := range words {
        tokens[word] = struct{}{}
    }
    return tokens
}

// jaccardSimilarity возвращает коэффициент Жаккара от 0.0 до 1.0
func jaccardSimilarity(a, b map[string]struct{}) float64 {
    intersection := 0
    union := make(map[string]struct{})
    
    for token := range a {
        union[token] = struct{}{}
        if _, exists := b[token]; exists {
            intersection++
        }
    }
    
    for token := range b {
        union[token] = struct{}{}
    }
    
    return float64(intersection) / float64(len(union))
}

// IsDuplicateFact возвращает true, если факты совпадают на более чем 70%
func (f Fact_Matcher) IsDuplicateFact(fact1, fact2 string) bool {
    norm1 := f.Normalize(fact1)
    norm2 := f.Normalize(fact2)
    
    tokens1 := tokenize(norm1)
    tokens2 := tokenize(norm2)
    
    similarity := jaccardSimilarity(tokens1, tokens2)
    
    return similarity >= 0.7
}
