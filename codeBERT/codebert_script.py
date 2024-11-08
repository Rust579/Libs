import sys
from transformers import AutoModelForSequenceClassification, AutoTokenizer
import torch

def main():
    code = sys.stdin.read().strip()
    # Пример кода и комментария
    # Код должен быть разделен с комментарием, например, `"<code>\n<comment>"`.
    code_part, comment_part = code.split('\n', 1)

    tokenizer = AutoTokenizer.from_pretrained("microsoft/codebert-base")
    model = AutoModelForSequenceClassification.from_pretrained("microsoft/codebert-base")

    # Объединение кода и комментария в одну строку
    inputs = tokenizer(f"{code_part}\n{comment_part}", return_tensors="pt", truncation=True, padding=True)
    outputs = model(**inputs)

    logits = outputs.logits
    probabilities = torch.nn.functional.softmax(logits, dim=-1)

    predicted_class = torch.argmax(probabilities, dim=-1)
    predicted_prob = probabilities[0][predicted_class].item()

    # Применение маппинга классов
    class_mapping = {
        0: "Not Relevant",
        1: "Relevant"
    }
    predicted_label = class_mapping.get(predicted_class.item(), "Unknown")

    print(f"Predicted class: {predicted_label}")
    print(f"Probability: {predicted_prob:.4f}")

if __name__ == "__main__":
    main()
